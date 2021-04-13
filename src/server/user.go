package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/CPEN391-Team-4/backend/src/imagestore"

	"github.com/CPEN391-Team-4/backend/src/logging"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


// Table for user storage
const USERS_TABLE = "users"

// Size for bufferr when reading data
const READ_BUF_SIZE = 4086

// User representation in database
type User struct {
	name       string
	image_id   string
	restricted bool
}

/********************* DATABASE *********************/

// addUserToDB Add a user to the USERS_TABLE
func (rs *routeServer) addUserToDB(name string, image_id string, restricted bool) error {
	restrict_int := 0
	if restricted {
		restrict_int = 1
	}
	sql := fmt.Sprintf(
		"INSERT INTO `%s` VALUES ('%s', '%s', '%d');",
		USERS_TABLE, name, image_id, restrict_int)
	_, err := rs.conn.Exec(sql)
	return err
}

// updateUserInDB Add a user in the USERS_TABLE
func (rs *routeServer) updateUserInDB(name string, image_id string, restricted bool) error {
	restrict_int := 0
	if restricted {
		restrict_int = 1
	}
	sql := fmt.Sprintf(
		"UPDATE `%s` SET image_id = '%s', restricted = '%d' WHERE name = '%s';",
		USERS_TABLE, image_id, restrict_int, name)
	fmt.Println(sql)
	_, err := rs.conn.Exec(sql)
	return err
}

// getAllUsersFromDB Retrieve all users in the USERS_TABLE
func (rs *routeServer) getAllUsersFromDB() ([]User, error) {
	sql := "SELECT * FROM " + USERS_TABLE
	results, err := rs.conn.Query(sql)

	var users []User
	users = make([]User, 0)

	if err != nil {
		return nil, err
	}
	for results.Next() {
		var u User
		err = results.Scan(&u.name, &u.image_id, &u.restricted)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
		fmt.Println("user=", u.name, "imageid=", u.image_id)
	}
	return users, nil
}

// getAllUserNameFromDB Retrieve all user names in the USERS_TABLE
func (rs *routeServer) getAllUserNameFromDB() ([]string, error) {
	var userNames []string
	userNames = make([]string, 0)

	users, err := rs.getAllUsersFromDB()

	if err != nil {
		return nil, err
	}
	for _, u := range users {
		userNames = append(userNames, u.name)
	}
	return userNames, nil
}

// getUserFromDB Retrieve a specific user from the USERS_TABLE
func (rs *routeServer) getUserFromDB(user string) (User, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE name = '%s';", USERS_TABLE, user)
	results, err := rs.conn.Query(sql)

	if err != nil {
		return User{}, err
	}

	for results.Next() {
		var u User
		err = results.Scan(&u.name, &u.image_id, &u.restricted)
		if err != nil {
			return User{}, err
		}
		return u, nil
	}

	return User{}, status.Errorf(codes.Unknown, "No user %s found", user)
}

// removeUserInDB Remove a specific user from the USERS_TABLE
func (rs *routeServer) removeUserInDB(name string) error {
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE name = '%s';", USERS_TABLE, name)
	_, err := rs.conn.Exec(sql)
	return err
}

/********************* USER API CALLS *********************/

// AddTrustedUser Receives a user, along with a stream of image data to be associated with them
func (rs *routeServer) AddTrustedUser(stream pb.Route_AddTrustedUserServer) error {
	imgBytes := bytes.Buffer{}
	var user *pb.User
	imageSize := 0
	for chunkNum := 0; ; chunkNum++ {

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logging.LogError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		if chunkNum == 0 {
			if req == nil {
				return logging.LogError(status.Errorf(codes.Unknown, "User must be set on first request"))
			}
			user = req
			log.Print("received a user: ", user.Name)
		}

		photo := req.GetPhoto()
		if photo != nil {
			chunk := photo.GetImage()
			size := len(chunk)

			_, err = imgBytes.Write(chunk)
			if err != nil {
				return logging.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			imageSize += size
		}

	}

	var id string
	var err error

	if imageSize != 0 {
		fw := imagestore.FileWriter{Directory: rs.imagestore}
		id, err = fw.Save("."+user.GetPhoto().FileExtension, imgBytes)
		if err != nil {
			return err
		}
	}

	err = rs.addUserToDB(user.GetName(), id, user.GetRestricted())
	if err != nil {
		return err
	}
	return stream.SendAndClose(&pb.Empty{})
}

// AddTrustedUser Receives an existing user to update, along with a stream of image data
// to be associated with them
func (rs *routeServer) UpdateTrustedUser(stream pb.Route_UpdateTrustedUserServer) error {
	imgBytes := bytes.Buffer{}
	var user *pb.User
	imageSize := 0
	for chunkNum := 0; ; chunkNum++ {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logging.LogError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		if chunkNum == 0 {
			if req == nil {
				return logging.LogError(status.Errorf(codes.Unknown, "User must be set on first request"))
			}
			user = req
			log.Print("received a user: ", user.Name)
		}

		photo := req.GetPhoto()
		if photo != nil {
			chunk := photo.GetImage()
			size := len(chunk)
			_, err = imgBytes.Write(chunk)
			if err != nil {
				return logging.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			imageSize += size
		}

	}

	var idUpdate string
	if imageSize > 0 {
		fw := imagestore.FileWriter{Directory: rs.imagestore}
		id, err := fw.Save("."+user.GetPhoto().FileExtension, imgBytes)
		if err != nil {
			return logging.LogError(status.Errorf(codes.Internal, "Failed saving image to disk: %v", err))
		}
		idUpdate = id
	}
	u, err := rs.getUserFromDB(user.GetName())
	if err != nil {
		return err
	}
	if len(u.image_id) > 0 {
		fw := imagestore.FileWriter{Directory: rs.imagestore}
		err = fw.Remove(u.image_id)
		if err != nil {
			return nil
		}
	}
	err = rs.updateUserInDB(user.GetName(), idUpdate, user.GetRestricted())
	if err != nil {
		return err
	}
	return stream.SendAndClose(&pb.Empty{})
}

// GetAllUserNames Retrieves all usernames from the database and returns the list to the client
func (rs *routeServer) GetAllUserNames(context.Context, *pb.Empty) (*pb.UserNames, error) {
	allUserNames, err := rs.getAllUserNameFromDB()
	if err != nil {
		return nil, logging.LogError(err)
	}
	return &pb.UserNames{
		Usernames: allUserNames,
	}, nil
}

// RemoveTrustedUser Removes a specific user from the database, as well as any images
func (rs *routeServer) RemoveTrustedUser(ctx context.Context, user *pb.User) (*pb.Empty, error) {
	u, err := rs.getUserFromDB(user.GetName())
	if err != nil {
		return &pb.Empty{}, err
	}
	err = rs.removeUserInDB(user.GetName())
	if err != nil {
		return &pb.Empty{}, err
	}

	if len(u.image_id) > 0 {
		fw := imagestore.FileWriter{Directory: rs.imagestore}
		err = fw.Remove(u.image_id)
	}

	return &pb.Empty{}, err
}

// GetUserPhoto Streams the photo for a specific user to the client
func (rs *routeServer) GetUserPhoto(user *pb.User, stream pb.Route_GetUserPhotoServer) error {
	if len(user.GetName()) == 0 {
		return status.Errorf(codes.Unknown, "User name not provided")
	}
	u, err := rs.getUserFromDB(user.GetName())
	if err != nil {
		return err
	}

	if len(u.image_id) == 0 {
		return nil
	}

	f, err := os.Open(rs.imagestore + "/" + u.image_id)
	if err != nil {
		return err
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, READ_BUF_SIZE)

	var photo pb.Photo

	sizeTotal := 0
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		photo.Image = buf[0:n]
		err = stream.Send(&photo)
		if err != nil {
			return err
		}
		sizeTotal += n
	}

	return nil
}
