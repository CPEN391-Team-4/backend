package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"github.com/CPEN391-Team-4/backend/src/logging"
	"github.com/CPEN391-Team-4/backend/src/videostore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const DEFAULT_ID = "default"
const VIDEOSTREAM_SIZE = 16
const TIME_INTERVAL = 3

type VideoStreams struct {
	sync.Mutex
	stream map[string]chan Frame
}

func (rs *routeServer) StreamVideo(stream pb.VideoRoute_StreamVideoServer) error {

	imgBytes := bytes.Buffer{}
	imageSize := 0
	created := false
	fw := videostore.FileWriter{Directory: rs.videostore}
	startFrame := true
	var dirId string

	rs.streams.Lock()
	if val, ok := rs.streams.stream[DEFAULT_ID]; !ok || rs.streams.stream[DEFAULT_ID] == nil {
		if val != nil {
			rs.streams.Unlock()
			return status.Errorf(codes.Unknown, "Stream id=%s is already live", DEFAULT_ID)
		}
		rs.streams.stream[DEFAULT_ID] = make(chan Frame, VIDEOSTREAM_SIZE)
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

			log.Printf("Videostream channel size=%v", len(rs.streams.stream[DEFAULT_ID]))
			_, err = fw.Save(dirId, int(frameNumber), imgBytes)
			if err != nil {
				return logging.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			// Remove oldest frame
			rs.streams.Lock()
			if len(rs.streams.stream[DEFAULT_ID]) >= VIDEOSTREAM_SIZE {
				<-rs.streams.stream[DEFAULT_ID]
			}
			rs.streams.stream[DEFAULT_ID] <- Frame{
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
	rs.streams.stream[DEFAULT_ID] = nil
	rs.streams.Unlock()

	return stream.SendAndClose(&pb.EmptyVideoResponse{})
}

func (rs *routeServer) PullVideoStream(req *pb.PullVideoStreamReq, stream pb.VideoRoute_PullVideoStreamServer) error {
	fmt.Println("Start live stream request received.")
	rs.video_stream_request = true
	time.Sleep(5 * time.Second)

	rs.streams.Lock()
	val, ok := rs.streams.stream[DEFAULT_ID]
	if !ok {
		rs.streams.Unlock()
		return status.Errorf(codes.Unknown, "Stream id=%s doesn't exist", DEFAULT_ID)
	}
	if val == nil {
		rs.streams.Unlock()
		return status.Errorf(codes.Unknown, "Stream id=%s is not live", DEFAULT_ID)
	}
	rs.streams.Unlock()
	for {
		rs.streams.Lock()

		val, ok := rs.streams.stream[DEFAULT_ID]
		if !ok || val == nil {
			err := stream.Send(&pb.PullVideoStreamResp{
				Closed: true,
			})
			rs.streams.Unlock()
			return err
		}
		if len(rs.streams.stream[DEFAULT_ID]) == 0 {
			rs.streams.Unlock()
			continue
		}
		f := <-rs.streams.stream[DEFAULT_ID]
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
	rs.video_stream_request = false
	return &pb.EmptyVideoResponse{}, nil
}

//keep sending the video_stream_request state to de1
func (rs *routeServer) RequestToStream(request *pb.InitialConnection, stream pb.VideoRoute_RequestToStreamServer) error {
	if request.Setup == false {

		return status.Errorf(codes.Unknown, "Did not set up the connection.", DEFAULT_ID)
	}
	fmt.Println("Successfully set up the connection with de1.")
	for {
		time.Sleep(TIME_INTERVAL * time.Second)
		stream.Send(&pb.Streamrequest{Request: rs.video_stream_request})
	}

	return status.Errorf(codes.Unknown, "Connection is broken.", DEFAULT_ID)
}
