package main

import (
	"fmt"
	"log"
	"net"
	"sistem-microservice/author/constants"
	"sistem-microservice/author/controllers"
	"sistem-microservice/author/models"
	authorpb "sistem-microservice/author/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := models.GetSqlConnection()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	fmt.Println("Connected to database")

	listener, err := net.Listen("tcp", ":"+constants.PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authorpb.RegisterAuthorServiceServer(grpcServer, controllers.NewAuthorController(db))

	fmt.Printf("Author service is running on port %s\n", constants.PORT)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
