package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/CPEN391-Team-4/backend/src/logging"
	"io"
	"log"
	"os"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const USERS_TABLE = "users"
const READ_BUF_SIZE = 1024

type User struct {
	name string
	image_id string
	restricted bool
}

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
	}
	return users, nil
}

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

func (rs *routeServer) removeUserInDB(name string) error {
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE name = '%s';", USERS_TABLE, name)
	_, err := rs.conn.Exec(sql)
	return err
}

func (rs *routeServer) AddTrustedUser(stream pb.Route_AddTrustedUserServer) error {
	// Referenced: dev.to/techschoolguru/
	//             upload-file-in-chunks-with-client-streaming-grpc-golang-4loc

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
			log.Print("received a user", user)
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
		fw := FileWriter{Directory: rs.imagestore}
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

func (rs *routeServer) UpdateTrustedUser(stream pb.Route_UpdateTrustedUserServer) error {
	// Referenced: dev.to/techschoolguru/
	//             upload-file-in-chunks-with-client-streaming-grpc-golang-4loc

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
			log.Print("received a user", user)
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
		fw := FileWriter{Directory: rs.imagestore}
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
		fw := FileWriter{Directory: rs.imagestore}
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

func (rs *routeServer) GetAllUserNames(context.Context, *pb.Empty) (*pb.UserNames, error) {
    fmt.Println("Hello")

	allUserNames, err := rs.getAllUserNameFromDB()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &pb.UserNames{
		Usernames: allUserNames,
	}, nil
}

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
		fw := FileWriter{Directory: rs.imagestore}
		err = fw.Remove(u.image_id)
	}

	return &pb.Empty{}, err
}

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
	fmt.Println("Sent %d bytes", sizeTotal)
	return nil
}