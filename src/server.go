package main

import (
	"database/sql"
	pb "github.com/CPEN391-Team-4/backend/pb/proto"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type routeServer struct {
	pb.UnimplementedRouteServer
	conn *sql.DB
	db string
	imagestore string
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

type env struct {
	db_uri string
	db string
	imagestore string
	server_address string
}

func (e *env) readEnv()  {
	e.db_uri = os.Getenv("DB_URI")
	e.db = os.Getenv("DB")
	e.imagestore = os.Getenv("IMAGESTORE")
	e.server_address = os.Getenv("SERVER_ADDRESS")
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

	rs := routeServer{conn: db, db: environ.db, imagestore: environ.imagestore}
	pb.RegisterRouteServer(grpcServer, &rs)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}