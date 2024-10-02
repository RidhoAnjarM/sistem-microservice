package main

import (
	"fmt"
	"log"
	"net"
	"github.com/raffa/book/constants"
	"github.com/raffa/book/controllers"
	"github.com/raffa/book/models"
	bookpb "github.com/raffa/book/proto"

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
	bookpb.RegisterBookServiceServer(grpcServer, controllers.NewBookController(db))

	fmt.Printf("Book service sedang berjalan di port %s\n", constants.PORT)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Gagal connect ke server: %v", err)
	}
}
