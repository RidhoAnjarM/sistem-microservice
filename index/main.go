package main

import (
	"fmt"
	"log"
	"net"
	"sistem-microservice/author/controllers"
	"sistem-microservice/author/models"
	authorpb "sistem-microservice/author/proto"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const(
	PORT = "8080"
	HTTP_PORT = "50053"
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

	// Set up HTTP server
	r := gin.Default()
	ac := controllers.NewAuthorController(db)

	// Daftarkan route
	r.GET("/authors/:id", ac.GetAuthorHandler) // Tambahkan ini

	go func() {
		listener, err := net.Listen("tcp", ":"+ PORT)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		authorpb.RegisterAuthorServiceServer(grpcServer, ac)

		fmt.Printf("Author service is running on port %s\n", PORT)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Mulai server HTTP
	if err := r.Run(":" + HTTP_PORT); err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
}
