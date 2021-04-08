package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/cognitiveservices/face"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"github.com/CPEN391-Team-4/backend/src/imagestore"
	"github.com/CPEN391-Team-4/backend/src/logging"
	"github.com/CPEN391-Team-4/backend/src/notification"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const userTimeout = 15

//const userToken = "dbW_Mb7ESuqlXw8lxY8YDs:APA91bHoKYX5YE-fWgzTP4uzJjW791Z7UlyW3tRsCoH33z8TUOHcHPRpZ5ZKy5bMXmCAvGzE3P7od9wt8R59X-rMmfDnpjNuVnBpGqqkRQ5J-m3VXcGHQMHdv2isw43-UfUAbBL5EP5y"

const userToken = "fzLRPJPvRHKvIBpKJb_uge:APA91bEDg0hxWa035bYDv2IjaR_PsDJMb8_QJHxefgMuPt-2c1t2EVLGKEt5QFt5hgpt4eKNMdXD-xzg6dYeFNcXG7VPIPIf_GjlgXR6_plosgbMHugZ8vbT6g0axvX-PwaaqBaYCYHC"

func (rs *routeServer) verifyFace(face0 *os.File, faceBuffer *bytes.Buffer) (*face.VerifyResult, error) {

	// A global context for use in all samples
	faceContext := context.Background()

	face0Closer := ioutil.NopCloser(face0)
	face1Closer := ioutil.NopCloser(faceBuffer)

	returnFaceIDVerify, returnFaceLandmarksVerify, returnRecognitionModelVerify := true, false, true

	// Detect face(s) from source image 1, returns a ListDetectedFace struct
	// We specify detection model 2 because we are not retrieving attributes.
	detectedVerifyFaces0, err := rs.faceClient.DetectWithStream(faceContext, face0Closer, &returnFaceIDVerify, &returnFaceLandmarksVerify, nil, face.Recognition03, &returnRecognitionModelVerify, face.Detection02)
	if err != nil {
		return nil, err
	}
	dVFaceIds0 := *detectedVerifyFaces0.Value
	imageSource0Id := dVFaceIds0[0].FaceID

	// Detect faces from each target image url in list. DetectWithURL returns a VerifyResult with Value of list[DetectedFaces]
	// Empty slice list for the target face IDs (UUIDs)
	var detectedVerifyFacesIds [2]uuid.UUID
	// We specify detection model 2 because we are not retrieving attributes.
	detectedVerifyFaces, err := rs.faceClient.DetectWithStream(faceContext, face1Closer, &returnFaceIDVerify, &returnFaceLandmarksVerify, nil, face.Recognition03, &returnRecognitionModelVerify, face.Detection02)
	if err != nil {
		return nil, err
	}

	dVFaces := *detectedVerifyFaces.Value
	// Add the returned face's face ID
	detectedVerifyFacesIds[0] = *dVFaces[0].FaceID

	// Verification example for faces of the same person. The higher the confidence, the more identical the faces in the images are.
	// Since target faces are the same person, in this example, we can use the 1st ID in the detectedVerifyFacesIds list to compare.
	verifyRequestBody1 := face.VerifyFaceToFaceRequest{FaceID1: imageSource0Id, FaceID2: &detectedVerifyFacesIds[0]}
	verifyResultSame, err := rs.faceClient.VerifyFaceToFace(faceContext, verifyRequestBody1)
	if err != nil {
		return nil, err
	}

	return &verifyResultSame, nil
}

type VerifyFaceResult struct {
	result *face.VerifyResult
	err    error
}

func (rs *routeServer) verifyFaceAsync(user *User, faceBuffer *bytes.Buffer) <-chan VerifyFaceResult {
	r := make(chan VerifyFaceResult)
	go func() {
		defer close(r)
		var res VerifyFaceResult
		if len(user.image_id) <= 0 {
			res.err = status.Errorf(codes.Unknown, "No image found for %s", user.name)
			res.result = nil
			r <- res
			return
		}
		faceOrig, err := os.Open(rs.imagestore + "/" + user.image_id)
		if err != nil {
			res.err = err
			r <- res
			return
		}

		res.result, res.err = rs.verifyFace(faceOrig, faceBuffer)
		if res.err != nil {
			r <- res
			return
		}
		res.err = faceOrig.Close()
		r <- res
	}()

	return r
}

func (rs *routeServer) VerifyUserFace(stream pb.Route_VerifyUserFaceServer) error {
	imgBytes := bytes.Buffer{}
	var fvReq *pb.FaceVerificationReq
	imageSize := 0
	for chunkNum := 0; ; chunkNum++ {

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logging.LogError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		if chunkNum == 0 {
			if req == nil {
				return logging.LogError(status.Errorf(codes.Unknown, "User must be set on first request"))
			}
			fvReq = req
			log.Print("received a request", fvReq)
		}

		photo := req.GetPhoto()
		if photo != nil {
			chunk := photo.GetImage()
			size := len(chunk)

			_, err = imgBytes.Write(chunk)
			if err != nil {
				return logging.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			imageSize += size
		}
	}
	var resp pb.FaceVerificationResp
	users, err := rs.getAllUsersFromDB()
	if err != nil {
		return err
	}
	resChan := make([]<-chan VerifyFaceResult, len(users))
	for i, user := range users {
		resChan[i] = rs.verifyFaceAsync(&user, &imgBytes)
	}

	dbuser := "Stranger"

	for i, user := range users {
		res := <-resChan[i]
		if res.err != nil {
			continue
		}
		if res.result == nil {
			continue
		}
		if *res.result.IsIdentical {
			fmt.Println(user.name)
			dbuser = user.name
			resp.User = user.name
			resp.Confidence = float32(*res.result.Confidence)
			break
		}
	}

	fmt.Println("Finishing verify face.")
	fw := imagestore.FileWriter{Directory: rs.imagestore}

	imgId, err := fw.Save(".raw", imgBytes)
	if err != nil {
		return logging.LogError(status.Errorf(codes.Internal, "cannot save image: %v", err))
	}

	recordID, err := rs.AddRecordToDB(dbuser, imgId)
	if err != nil {
		return logging.LogError(status.Errorf(codes.Internal, "cannot add record to db: %v", err))
	}

	// TODO: tok
	_, err = notification.Send(userToken, "Detected human motion", fmt.Sprintf("user=%s", resp.User), rs.firebaseKeyfile)
	if err != nil {
		return logging.LogError(status.Errorf(codes.Internal, "cannot send notification: %v", err))
	}

	// Wait for response
	select {
	case res := <-rs.waitingUser:
		fmt.Println("Success:", res)
		rs.UpdateRecordStatusToDB(recordID, "Agree")
		resp.Accept = res == permAllow
	case <-time.After(userTimeout * time.Second):
		log.Println("timeout waiting on notification response")
		rs.UpdateRecordStatusToDB(recordID, "Deny")
		resp.User = ""
		resp.Confidence = 0
		resp.Accept = false
	}

	return stream.SendAndClose(&resp)
}
