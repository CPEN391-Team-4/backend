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

type routeServer struct {
	pb.UnimplementedRouteServer
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

		chunk := req.GetImage().Image
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)
		log.Print(chunk)

		_, err = imgBytes.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}

		imageSize += size
	}

	log.Print(imgBytes)
	// TODO: Log to database
	fw := FileWriter{Directory: "./imagestore"}
	id, err := fw.Save("."+user.GetImage().FileExtension, imgBytes)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "Failed saving image to disk: %v", err))
	}

	fmt.Println(id)

	return nil

}

func (rs *routeServer) UpdateTrustedUser(stream pb.Route_UpdateTrustedUserServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateTrustedUser not implemented")
}
func (rs *routeServer) RemoveTrustedUser(context.Context, *pb.User) (*pb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTrustedUser not implemented")
}