package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sistem-microservice/author/constants"
	"sistem-microservice/author/controllers"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		log.Fatal("POSTGRES_URL not set in .env file")
	}

	db, err := pgxpool.Connect(context.Background(), postgresURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected to database")

	listener, err := net.Listen("tcp", ":"+constants.PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authorpb.RegisterAuthorServiceServer(grpcServer, controllers.NewAuthorController(db))

	fmt.Printf("Server is running on port %s\n", constants.PORT)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
