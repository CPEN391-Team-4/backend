package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/face"
	"github.com/Azure/go-autorest/autorest"
	"github.com/gofrs/uuid"
)

func main2() {

	// // A global context for use in all samples
	// faceContext := context.Background()

	// // Base url for the Verify and Large Face List examples
	// const imageBaseURL = "https://csdx.blob.core.windows.net/resources/Face/Images/"

	// /*
	//    Authenticate
	// */
	// // Add FACE_SUBSCRIPTION_KEY, FACE_ENDPOINT, and AZURE_SUBSCRIPTION_ID to your environment variables.
	// subscriptionKey := os.Getenv("FACE_SUBSCRIPTION_KEY")
	// endpoint := os.Getenv("FACE_ENDPOINT")

	// // Client used for Detect Faces, Find Similar, and Verify examples.
	// client := face.NewClient(endpoint)

	// fmt.Println(client)
	// client.Authorizer = autorest.NewCognitiveServicesAuthorizer(subscriptionKey)
	// /*
	//    END - Authenticate
	// */
	// // Detect a face in an image that contains a single face
	// singleFaceImageURL := "https://www.biography.com/.image/t_share/MTQ1MzAyNzYzOTgxNTE0NTEz/john-f-kennedy---mini-biography.jpg"
	// singleImageURL := face.ImageURL{URL: &singleFaceImageURL}
	// singleImageName := path.Base(singleFaceImageURL)
	// fmt.Println(singleImageName)
	// // Array types chosen for the attributes of Face
	// attributes := []face.AttributeType{"age", "emotion", "gender"}
	// returnFaceID := true
	// returnRecognitionModel := false
	// returnFaceLandmarks := false

	// // API call to detect faces in single-faced image, using recognition model 3
	// // We specify detection model 1 because we are retrieving attributes.
	// detectSingleFaces, dErr := client.DetectWithURL(faceContext, singleImageURL, &returnFaceID, &returnFaceLandmarks, attributes, face.Recognition03, &returnRecognitionModel, face.Detection01)
	// if dErr != nil {
	// 	log.Fatal(dErr)
	// }

	// // Dereference *[]DetectedFace, in order to loop through it.
	// dFaces := *detectSingleFaces.Value
	// fmt.Println("Detected face in (" + singleImageName + ") with ID(s): ")
	// fmt.Println(dFaces[0].FaceID)
	// fmt.Println()

	// // Find/display the age and gender attributes
	// for _, dFace := range dFaces {
	// 	fmt.Println("Face attributes:")
	// 	fmt.Printf("  Age: %.0f", *dFace.FaceAttributes.Age)
	// 	fmt.Println("\n  Gender: " + dFace.FaceAttributes.Gender)
	// }
	verifyFace()
}

func verifyFace() {

	// A global context for use in all samples
	faceContext := context.Background()

	// Base url for the Verify and Large Face List examples
	const imageBaseURL = "https://csdx.blob.core.windows.net/resources/Face/Images/"

	/*
	   Authenticate
	*/
	// Add FACE_SUBSCRIPTION_KEY, FACE_ENDPOINT, and AZURE_SUBSCRIPTION_ID to your environment variables.
	subscriptionKey := os.Getenv("FACE_SUBSCRIPTION_KEY")
	endpoint := os.Getenv("FACE_ENDPOINT")

	// Client used for Detect Faces, Find Similar, and Verify examples.
	client := face.NewClient(endpoint)

	client.Authorizer = autorest.NewCognitiveServicesAuthorizer(subscriptionKey)
	/*
	   END - Authenticate
	*/

	faceimage1, err := os.Open("IMAGESTORE/face1.jpg")
	if err != nil {
		fmt.Println(err)
	}
	faceimage1ReadCloser1 := ioutil.NopCloser(faceimage1)

	faceimage2, err := os.Open("IMAGESTORE/face2.jpg")
	if err != nil {
		fmt.Println(err)
	}
	faceimage1ReadCloser2 := ioutil.NopCloser(faceimage2)

	//urlSource1 := "https://bostonglobe-prod.cdn.arcpublishing.com/resizer/jWdSsVXZIgPqW4nt5fchsdj3SHg=/1440x0/arc-anglerfish-arc2-prod-bostonglobe.s3.amazonaws.com/public/PS7FS3VGEYI6RAOLA4BXWMMOGA.jpg"
	//url1 := face.ImageURL{URL: &urlSource1}

	returnFaceIDVerify := true
	returnFaceLandmarksVerify := false
	returnRecognitionModelVerify := false

	// Detect face(s) from source image 1, returns a ListDetectedFace struct
	// We specify detection model 2 because we are not retrieving attributes.
	detectedVerifyFaces1, dErrV1 := client.DetectWithStream(faceContext, faceimage1ReadCloser1, &returnFaceIDVerify, &returnFaceLandmarksVerify, nil, face.Recognition03, &returnRecognitionModelVerify, face.Detection02)
	if dErrV1 != nil {
		log.Fatal(dErrV1)
	}
	// Dereference the result, before getting the ID
	dVFaceIds1 := *detectedVerifyFaces1.Value
	// Get ID of the detected face
	imageSource1Id := dVFaceIds1[0].FaceID
	fmt.Println(fmt.Sprintf("%v face(s) detected from image: %v", len(dVFaceIds1), "face1.jpg"))

	// Detect faces from each target image url in list. DetectWithURL returns a VerifyResult with Value of list[DetectedFaces]
	// Empty slice list for the target face IDs (UUIDs)
	var detectedVerifyFacesIds [2]uuid.UUID

	// urlSource := "https://cdn.cnn.com/cnnnext/dam/assets/151216201256-faces-of-donald-trump-sequel-moos-dnt-erin-00013313-exlarge-169.jpg"
	// url := face.ImageURL{URL: &urlSource}
	// We specify detection model 2 because we are not retrieving attributes.
	detectedVerifyFaces, dErrV := client.DetectWithStream(faceContext, faceimage1ReadCloser2, &returnFaceIDVerify, &returnFaceLandmarksVerify, nil, face.Recognition03, &returnRecognitionModelVerify, face.Detection02)
	if dErrV != nil {
		log.Fatal(dErrV)
	}
	// Dereference *[]DetectedFace from Value in order to loop through it.
	dVFaces := *detectedVerifyFaces.Value
	// Add the returned face's face ID
	detectedVerifyFacesIds[0] = *dVFaces[0].FaceID
	fmt.Println(fmt.Sprintf("%v face(s) detected from image: %v", len(dVFaces), "face2.jpg"))

	// Verification example for faces of the same person. The higher the confidence, the more identical the faces in the images are.
	// Since target faces are the same person, in this example, we can use the 1st ID in the detectedVerifyFacesIds list to compare.
	verifyRequestBody1 := face.VerifyFaceToFaceRequest{FaceID1: imageSource1Id, FaceID2: &detectedVerifyFacesIds[0]}
	verifyResultSame, vErrSame := client.VerifyFaceToFace(faceContext, verifyRequestBody1)
	if vErrSame != nil {
		log.Fatal(vErrSame)
	}

	fmt.Println()

	// Check if the faces are from the same person.
	if *verifyResultSame.IsIdentical {
		fmt.Println(fmt.Sprintf(" same person,  confidence %v",
			strconv.FormatFloat(*verifyResultSame.Confidence, 'f', 3, 64)))
	} else {
		// Low confidence means they are more differant than same.
		fmt.Println(fmt.Sprintf(" different person, with confidence %v",
			strconv.FormatFloat(*verifyResultSame.Confidence, 'f', 3, 64)))
	}

}
