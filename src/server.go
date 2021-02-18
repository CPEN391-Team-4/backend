package main

import (
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
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

	var user *pb.User
	imageSize := 0
	for chunkNum := 0;; chunkNum++ {
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
			user = req
			log.Print("received a user", user)
		}

		chunk := req.GetImage().Image
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)
		log.Print(chunk)

		imageSize += size
	}

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