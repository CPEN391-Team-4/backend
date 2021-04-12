package main

import (
	"database/sql"
	"log"
	"net"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	"github.com/CPEN391-Team-4/backend/src/environment"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/face"
	"github.com/Azure/go-autorest/autorest"
)

type Frame struct {
	number    int
	data      []byte
	lastChunk bool
}

const (
	permWaiting = 0x01
	permAllow   = 0x02
	permDeny    = 0x03
)

type routeServer struct {
	pb.UnimplementedRouteServer
	pb.UnimplementedVideoRouteServer
	conn               *sql.DB
	db                 string
	imagestore         string
	videostore         string
	faceClient         *face.Client
	streams            VideoStreams
	firebaseKeyfile    string
	videoStreamRequest VideoStreamRequest
	unlockDoorRequest  UnlockDoorRequest
}

func main() {
	environ := environment.Env{}
	environ.ReadEnv()
	log.Println("Listening on:", environ.ServerAddress)

	lis, err := net.Listen("tcp", environ.ServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := sql.Open("mysql", environ.DbUri)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer db.Close()
	defer lis.Close()

	grpcServer := grpc.NewServer()

	// Client used for Detect Faces, Find Similar, and Verify examples.
	faceClient := face.NewClient(environ.FaceEndpoint)
	faceClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(environ.FaceSubscriptionKey)

	rs := routeServer{
		conn:            db,
		db:              environ.Db,
		imagestore:      environ.Imagestore,
		videostore:      environ.Videostore,
		firebaseKeyfile: environ.FirebaseKeyfile,
		faceClient:      &faceClient,
		streams:         VideoStreams{stream: make(map[string]chan Frame)},
		videoStreamRequest: VideoStreamRequest{
			requested: make(chan bool, 1),
			up:        make(chan bool, 1),
		},
		unlockDoorRequest: UnlockDoorRequest{
			requested: make(chan bool, 1),
			done:      make(chan bool, 1),
		},
	}
	pb.RegisterRouteServer(grpcServer, &rs)
	pb.RegisterVideoRouteServer(grpcServer, &rs)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
