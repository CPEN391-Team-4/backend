package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

		photo := req.GetImage()
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

	fmt.Println(imageSize)

	return nil

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
