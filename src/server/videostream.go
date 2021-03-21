package main

import (
	"bytes"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"github.com/CPEN391-Team-4/backend/src/logging"
	"github.com/CPEN391-Team-4/backend/src/videostore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

func (rs *routeServer) StreamVideo(stream pb.VideoRoute_StreamVideoServer) error {

	imgBytes := bytes.Buffer{}
	imageSize := 0
	created := false
	fw := videostore.FileWriter{Directory: rs.videostore}
	startFrame := true
	var dirId string
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

			_, err = imgBytes.Write(chunk)
			if err != nil {
				return logging.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			imageSize += size

			_, err = fw.Save(dirId, int(frameNumber), imgBytes)
			if err != nil {
				return logging.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			if lastChunk {
				startFrame = true
				imgBytes = bytes.Buffer{}
			}
		}

	}

	return stream.SendAndClose(&pb.EmptyVideoResponse{})
}