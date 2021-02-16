package main

import (
	"context"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
	"github.com/CPEN391-Team-4/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type routeServer struct {
	pb.UnimplementedRouteServer
}

func (s *routeServer) GetUsers(context.Context, *pb.UsersParams) (*pb.Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}