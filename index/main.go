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
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	db, err := models.GetSqlConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Connected to database")

	r := gin.Default()
	ac := controllers.NewAuthorController(db)

	// Register HTTP routes
	r.POST("/authors/create", ac.CreateAuthorHandler)  
	r.GET("/authors/:id", ac.GetAuthorHandler) 
	r.DELETE("/authors/:id", ac.DeleteAuthorHandler) 

	// Jalankan HTTP server di goroutine
	go func() {
		if err := r.Run(":" + HTTP_PORT); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()

	// Set up gRPC server
	listener, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", HTTP_PORT, err)
	}

	grpcServer := grpc.NewServer()
	authorpb.RegisterAuthorServiceServer(grpcServer, ac) 

	fmt.Printf("gRPC server is running on port %s\n", HTTP_PORT)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
