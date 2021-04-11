package main

import (
	"context"
	"fmt"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"io"
)

type UnlockDoorRequest struct {
	requested chan bool
	done chan bool
}

func (rs *routeServer) LockDoor(ctx context.Context, req *pb.LockDoorReq) (*pb.LockResp, error) {
	fmt.Printf("Lock request received: %v", req)
	rs.videoStreamRequest.requested <- req.Locked
	return &pb.LockResp{Success: <- rs.unlockDoorRequest.done}, nil
}

//keep sending the video_stream_request state to de1
func (rs *routeServer) RequestToLock(stream pb.Route_RequestToLockServer) error {
	for {
		req := <- rs.unlockDoorRequest.requested
		err := stream.Send(&pb.LockReq{Request: req})
		if err != nil {
			return err
		}
		in, err := stream.Recv()
		if err != nil {
			return err
		}
		if err == io.EOF {
			return nil
		}
		rs.unlockDoorRequest.done <- in.Setup
	}

	return nil
}