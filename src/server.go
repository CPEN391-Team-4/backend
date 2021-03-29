package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	pb "github.com/CPEN391-Team-4/backend/pb/proto"
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

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

type env struct {
	db_uri                string
	db                    string
	imagestore            string
	server_address        string
	face_subscription_key string
	face_endpoint         string
}

func (e *env) readEnv() {
	e.db_uri = os.Getenv("DB_URI")
	e.db = os.Getenv("DB")
	e.imagestore = os.Getenv("IMAGESTORE")
	e.server_address = os.Getenv("SERVER_ADDRESS")
	e.face_subscription_key = os.Getenv("FACE_SUBSCRIPTION_KEY")
	e.face_endpoint = os.Getenv("FACE_ENDPOINT")
}

func main() {
	environ := env{}
	environ.readEnv()

	lis, err := net.Listen("tcp", environ.server_address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := sql.Open("mysql", environ.db_uri)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()

	// Client used for Detect Faces, Find Similar, and Verify examples.
	faceClient := face.NewClient(environ.face_endpoint)
	faceClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(environ.face_subscription_key)

	rs := routeServer{
		conn:       db,
		db:         environ.db,
		imagestore: environ.imagestore,
		faceClient: &faceClient,
	}
	pb.RegisterRouteServer(grpcServer, &rs)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	//rs.test_history()

}
