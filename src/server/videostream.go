package main

import (
	"bytes"
	"context"
	"fmt"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"github.com/CPEN391-Team-4/backend/src/logging"
	"github.com/CPEN391-Team-4/backend/src/videostore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"sync"
)

const VIDEOSTREAM_SIZE = 16
const TIME_INTERVAL = 5

type VideoStreams struct {
	sync.Mutex
	stream map[string]chan Frame
}
type VideoStreamRequest struct {
	requested chan bool
	up        chan bool
}

func (rs *routeServer) StreamVideo(stream pb.VideoRoute_StreamVideoServer) error {

	imgBytes := bytes.Buffer{}
	imageSize := 0
	created := false
	fw := videostore.FileWriter{Directory: rs.videostore}
	startFrame := true
	var dirId string

	req, err := stream.Recv()
	if err != nil {
		return logging.LogError(status.Errorf(codes.Unknown, "cannot receive initial data: %v", err))
	}

	did := req.DeviceId

	rs.streams.Lock()
	if val, ok := rs.streams.stream[did]; !ok || rs.streams.stream[did] == nil {
		if val != nil {
			rs.streams.Unlock()
			return status.Errorf(codes.Unknown, "Stream id=%s is already live", did)
		}
		rs.streams.stream[did] = make(chan Frame, VIDEOSTREAM_SIZE)
	}
	rs.streams.Unlock()

	for chunkNum := 0; ; chunkNum++ {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logging.LogError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		frame := req.GetFrame()
		if frame != nil {
			chunk := frame.GetChunk()
			lastChunk := frame.GetLastChunk()
			frameNumber := frame.GetNumber()
			size := len(chunk)

			if !created && startFrame && frameNumber == 0 {
				dirId, err = fw.CreateSubdir()
				if err != nil {
					return logging.LogError(err)
				}
				startFrame = false
				created = true
			}

			log.Printf("Recieved frame=%v, lastChunk=%v", frameNumber, lastChunk)

			_, err = imgBytes.Write(chunk)
			if err != nil {
				return logging.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			imageSize += size

			if !lastChunk {
				continue
			}

			log.Printf("Videostream channel size=%v", len(rs.streams.stream[did]))
			_, err = fw.Save(dirId, int(frameNumber), imgBytes)
			if err != nil {
				return logging.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			// Remove oldest frame
			rs.streams.Lock()
			if len(rs.streams.stream[did]) >= VIDEOSTREAM_SIZE {
				<-rs.streams.stream[did]
			}
			rs.streams.stream[did] <- Frame{
				number:    int(frameNumber),
				data:      imgBytes.Bytes(),
				lastChunk: lastChunk,
			}
			rs.streams.Unlock()

			startFrame = true
			imgBytes = bytes.Buffer{}
		}
	}
	rs.streams.Lock()
	rs.streams.stream[did] = nil
	rs.streams.Unlock()

	return stream.SendAndClose(&pb.EmptyVideoResponse{})
}

func (rs *routeServer) PullVideoStream(req *pb.PullVideoStreamReq, stream pb.VideoRoute_PullVideoStreamServer) error {
	fmt.Println("Start live stream request received.")
	rs.videoStreamRequest.requested <- true

	// Wait for 'up'
	<-rs.videoStreamRequest.up

	did, err := rs.getDe1IDFromDB(req.MainUser)
	if err != nil {
		return err
	}

	rs.streams.Lock()
	val, ok := rs.streams.stream[did]
	if !ok {
		rs.streams.Unlock()
		return status.Errorf(codes.Unknown, "Stream id=%s doesn't exist", did)
	}
	if val == nil {
		rs.streams.Unlock()
		return status.Errorf(codes.Unknown, "Stream id=%s is not live", did)
	}
	rs.streams.Unlock()
	for {
		rs.streams.Lock()

		val, ok := rs.streams.stream[did]
		if !ok || val == nil {
			err := stream.Send(&pb.PullVideoStreamResp{
				Closed: true,
			})
			rs.streams.Unlock()
			return err
		}
		if len(rs.streams.stream[did]) == 0 {
			rs.streams.Unlock()
			continue
		}
		f := <-rs.streams.stream[did]
		err := stream.Send(&pb.PullVideoStreamResp{
			Video: &pb.Video{
				Frame: &pb.Frame{
					Number:    int32(f.number),
					LastChunk: f.lastChunk,
					Chunk:     f.data,
				},
			},
			Closed: false,
		})
		rs.streams.Unlock()
		if err != nil {
			return err
		}
	}

	return nil
}

// receive call from app to end the stream
func (rs *routeServer) EndPullVideoStream(ctx context.Context, request *pb.EndPullVideoStreamReq) (*pb.EmptyVideoResponse, error) {
	fmt.Println("End live stream request received.")
	rs.videoStreamRequest.requested <- false
	return &pb.EmptyVideoResponse{}, nil
}

//keep sending the video_stream_request state to de1
func (rs *routeServer) RequestToStream(stream pb.VideoRoute_RequestToStreamServer) error {
	for {
		var req bool
		select {
		case <-stream.Context().Done():
			log.Println("done RequestToStream")
			return nil
		case req = <-rs.videoStreamRequest.requested:
			break
		}
		err := stream.Send(&pb.Streamrequest{Request: req})
		if err != nil {
			return err
		}

		if req {
			in, err := stream.Recv()
			log.Println("receive up from backend")
			if err != nil {
				return err
			}
			if err == io.EOF {
				return nil
			}
			rs.videoStreamRequest.up <- in.Setup
		}
	}

	return nil
}
