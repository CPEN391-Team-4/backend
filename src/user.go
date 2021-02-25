package main

import (
	"bytes"
	"context"
	"fmt"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

const USERS_TABLE = "users"

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


func (rs *routeServer) updateUserInDB(name string, image_id *string, restricted bool) error {
	restrict_int := 0
	if restricted {
		restrict_int = 1
	}
	sql := "UPDATE `" + USERS_TABLE + "` SET "
	if image_id != nil {
		sql += "image_id = '" + *image_id + "', "
	}
	sql += fmt.Sprintf("restricted = '%d' WHERE name = '%s';", restrict_int, name)
	_, err := rs.conn.Exec(sql)
	return err
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
		log.Println("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		if chunkNum == 0 {
			if req == nil {
				return logError(status.Errorf(codes.Unknown, "User must be set on first request"))
			}
			user = req
			log.Print("received a user", user)
		}

		photo := req.GetPhoto()
		if photo != nil {
			chunk := photo.GetImage()
			size := len(chunk)

			log.Printf("received a chunk with size: %d", size)
			log.Print(chunk)

			_, err = imgBytes.Write(chunk)
			if err != nil {
				return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			imageSize += size
		}

	}

	fw := FileWriter{Directory: rs.imagestore}
	id, err := fw.Save("."+user.GetPhoto().FileExtension, imgBytes)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "Failed saving image to disk: %v", err))
	}

	fmt.Println(id)

	return rs.addUserToDB(user.GetName(), id, user.GetRestricted())

}

func (rs *routeServer) UpdateTrustedUser(stream pb.Route_UpdateTrustedUserServer) error {
	// Referenced: dev.to/techschoolguru/
	//             upload-file-in-chunks-with-client-streaming-grpc-golang-4loc

	imgBytes := bytes.Buffer{}
	var user *pb.User
	imageSize := 0
	for chunkNum := 0; ; chunkNum++ {
		log.Println("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		if chunkNum == 0 {
			if req == nil {
				return logError(status.Errorf(codes.Unknown, "User must be set on first request"))
			}
			user = req
			log.Print("received a user", user)
		}

		photo := req.GetPhoto()
		if photo != nil {
			chunk := photo.GetImage()
			size := len(chunk)

			log.Printf("received a chunk with size: %d", size)
			log.Print(chunk)

			_, err = imgBytes.Write(chunk)
			if err != nil {
				return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			imageSize += size
		}

	}

	var idUpdate *string = nil
	if imageSize > 0 {
		fw := FileWriter{Directory: rs.imagestore}
		id, err := fw.Save("."+user.GetPhoto().FileExtension, imgBytes)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "Failed saving image to disk: %v", err))
		}
		idUpdate = &id
	}


	return rs.updateUserInDB(user.GetName(), idUpdate, user.GetRestricted())

}

func (rs *routeServer) RemoveTrustedUser(ctx context.Context, user *pb.User) (*pb.Empty, error) {
	err := rs.removeUserInDB(user.GetName())
	return &pb.Empty{}, err
}