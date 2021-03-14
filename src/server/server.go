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

type routeServer struct {
	pb.UnimplementedRouteServer
	conn       *sql.DB
	db         string
	imagestore string
	faceClient *face.Client
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
		faceClient: &faceClient,
	}
	pb.RegisterRouteServer(grpcServer, &rs)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
