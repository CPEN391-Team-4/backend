package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"google.golang.org/grpc"
)

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

type routeServer struct {
	pb.UnimplementedRouteServer
}

func (rs *routeServer) AddTrustedUser(stream pb.Route_AddTrustedUserServer) (*pb.Response, error) {
	// Referenced: dev.to/techschoolguru/
	//             upload-file-in-chunks-with-client-streaming-grpc-golang-4loc
	response := &pb.Response{}
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
			return response, nil
		}

		if chunkNum == 0 {
			if req == nil {
				return response, nil
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
			return response, nil
		}

		imageSize += size
	}

	log.Print(imgBytes)
	// TODO: Log to database
	// fw := FileWriter{Directory: "./imagestore"}
	// id, err := fw.Save("." + user.GetImage().FileExtension, imgBytes)
	// if err != nil {
	// 	return logError(status.Errorf(codes.Internal, "Failed saving image to disk: %v", err))
	// }

	//fmt.Println(id)

	return response, nil
}

func (rs *routeServer) RemoveTrustedUser(ctx context.Context, u *pb.User) pb.Response {

	fmt.Println("Recieved", u)
	response := pb.Response{}
	return response
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	rs := routeServer{}
	pb.RegisterRouteServer(grpcServer, &rs)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
