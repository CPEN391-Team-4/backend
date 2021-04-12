package main

import (
	"context"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"io"
	"log"
)

type UnlockDoorRequest struct {
	requested chan bool
	done      chan bool
}

func (rs *routeServer) LockDoor(ctx context.Context, req *pb.LockDoorReq) (*pb.LockResp, error) {
	log.Printf("Lock request received: %v", req)
	rs.unlockDoorRequest.requested <- req.Locked
	log.Printf("Lock request fin: %v", req)
	log.Printf("LockDoor: len(rs.unlockDoorRequest.requested)=%v", len(rs.unlockDoorRequest.requested))
	return &pb.LockResp{Success: <-rs.unlockDoorRequest.done}, nil
}

//keep sending the video_stream_request state to de1
func (rs *routeServer) RequestToLock(stream pb.Route_RequestToLockServer) error {
	for {
		log.Printf("RequestToLock")
		log.Printf("ReqToLock: len(rs.unlockDoorRequest.requested)=%v", len(rs.unlockDoorRequest.requested))

		var req bool
		select {
		case <-stream.Context().Done():
			log.Println("done RequestToLock")
			return nil
		case req = <-rs.unlockDoorRequest.requested:
			break
		}
		log.Printf("ReqToLock: len(rs.unlockDoorRequest.requested)=%v", len(rs.unlockDoorRequest.requested))
		log.Printf("req=%v", req)
		err := stream.Send(&pb.LockReq{Request: req})
		if err != nil {
			return err
		}

		log.Printf("Sent=%v", req)
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
