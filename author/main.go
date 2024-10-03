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

	// Membuat koneksi database menggunakan GORM
	db, err := models.GetSqlConnection()
	if err != nil {
		log.Fatalf("gagal connect ke database: %v", err)
	}

	fmt.Println("berhasil connect ke database")

	// Set up gRPC server
	listener, err := net.Listen("tcp", ":"+constants.PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authorpb.RegisterAuthorServiceServer(grpcServer, controllers.NewAuthorController(db))

	fmt.Printf("Author service sedang berjalan di port %s\n", constants.PORT)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Gagal connect ke server: %v", err)
	}
}
