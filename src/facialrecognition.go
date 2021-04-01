package main

import (
	"bytes"
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/cognitiveservices/face"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"log"
	"os"
)

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
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		if chunkNum == 0 {
			if req == nil {
				return logError(status.Errorf(codes.Unknown, "User must be set on first request"))
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
				return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}

			imageSize += size
		}
	}
	user := fvReq.GetUser()
	u, err := rs.getUserFromDB(user)
	if err != nil {
		return err
	}
	if len(u.image_id) <= 0 {
		return status.Errorf(codes.Unknown, "No image found for %s", u.name)
	}
	faceOrig, err := os.Open(rs.imagestore + "/" + u.image_id)
	if err != nil {
		return err
	}

// TESTING
//	imgBytes2, err := ioutil.ReadFile(rs.imagestore + "/" + u.image_id)
//	if err != nil {
//		return err
//	}
//	imgBytesBuffer := bytes.NewBuffer(imgBytes2)
//	verif, err := rs.verifyFace(faceOrig, imgBytesBuffer)
// END TESTING

	verif, err := rs.verifyFace(faceOrig, &imgBytes)

	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.FaceVerificationResp{Verified: *verif.IsIdentical})
}