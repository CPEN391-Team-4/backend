package main

import (
	"context"
	"fmt"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"google.golang.org/grpc"
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

func (rs *routeServer) AddTrustedUser(ctx context.Context, u *pb.User) (*pb.Empty, error) {

	fmt.Println("Recieved", u)

	return &pb.Empty{}, nil
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