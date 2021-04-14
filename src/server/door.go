package main

import (
	"context"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"io"
	"log"
)

// Flags for door state
type UnlockDoorRequest struct {
	requested chan bool
	done      chan bool
}

// LockDoor Set the lock/unlock door request and respond when complete
func (rs *routeServer) LockDoor(ctx context.Context, req *pb.LockDoorReq) (*pb.LockResp, error) {
	rs.unlockDoorRequest.requested <- req.Locked
	return &pb.LockResp{Success: <-rs.unlockDoorRequest.done}, nil
}

// RequestToLock Persistent connection to send lock state to client
func (rs *routeServer) RequestToLock(stream pb.Route_RequestToLockServer) error {
	for {
		var req bool
		select {
		case <-stream.Context().Done():
			return nil
		case req = <-rs.unlockDoorRequest.requested:
			break
		}
		log.Printf("req=%v", req)
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
		log.Printf("Recv=%v", in)
		rs.unlockDoorRequest.done <- in.Setup
	}

	return nil
}
