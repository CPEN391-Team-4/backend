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

type Frames struct {
	f Frame
	current int
}
type Frame struct {
	number int
	data []byte
}

type routeServer struct {
	pb.UnimplementedRouteServer
	pb.UnimplementedVideoRouteServer
	conn       *sql.DB
	db         string
	imagestore string
	videostore string
	faceClient *face.Client
	streams    map[string][]Frames
}


func main() {
	environ := environment.Env{}
	environ.ReadEnv()

	lis, err := net.Listen("tcp", environ.ServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := sql.Open("mysql", environ.DbUri)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()

	// Client used for Detect Faces, Find Similar, and Verify examples.
	faceClient := face.NewClient(environ.FaceEndpoint)
	faceClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(environ.FaceSubscriptionKey)

	rs := routeServer{
		conn: db,
		db: environ.Db,
		imagestore: environ.Imagestore,
		videostore: environ.Videostore,
		faceClient: &faceClient,
	}
	pb.RegisterRouteServer(grpcServer, &rs)
	pb.RegisterVideoRouteServer(grpcServer, &rs)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
