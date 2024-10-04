package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"sistem-microservice/author/constants"
	"sistem-microservice/author/controllers"
	"sistem-microservice/author/models"
	authorpb "sistem-microservice/author/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load file .env untuk konfigurasi
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Membuat koneksi ke database menggunakan GORM
	db, err := models.GetSqlConnection()
	if err != nil {
		log.Fatalf("Gagal connect ke database: %v", err)
	}

	fmt.Println("Berhasil connect ke database")

	// Membuat listener untuk server gRPC pada port yang ditentukan
	listener, err := net.Listen("tcp", ":"+constants.PORT)
	if err != nil {
		log.Fatalf("Gagal listen pada port %v: %v", constants.PORT, err)
	}

	// Membuat instance gRPC server
	grpcServer := grpc.NewServer()

	// Membuat instance AuthorController
	authorController := controllers.NewAuthorController(db)

	// Mendaftarkan AuthorController sebagai service gRPC
	authorpb.RegisterAuthorServiceServer(grpcServer, authorController)

	// Jalankan gRPC Server secara bersamaan dengan HTTP Server
	go func() {
		fmt.Printf("Author gRPC service sedang berjalan di port %s\n", constants.PORT)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Gagal menjalankan server gRPC: %v", err)
		}
	}()

	http.HandleFunc("/authors/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/authors/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid author ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			// Handler untuk mengambil author
			req := &authorpb.GetAuthorRequest{Id: int32(id)}
			res, err := authorController.GetAuthorWithBooks(context.Background(), req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			// Response sebagai JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)

		case http.MethodDelete:
			// Handler untuk menghapus author
			var author models.Author
			if err := authorController.DB.First(&author, id).Error; err != nil {
				http.Error(w, "Author not found", http.StatusNotFound)
				return
			}

			// Hapus buku-buku yang terkait
			if err := authorController.DB.Where("author_id = ?", id).Delete(&models.Book{}).Error; err != nil {
				http.Error(w, "Failed to delete related books", http.StatusInternalServerError)
				return
			}

			// Hapus author
			if err := authorController.DB.Delete(&author).Error; err != nil {
				http.Error(w, "Failed to delete author", http.StatusInternalServerError)
				return
			}

			// Response sukses
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Author and related books deleted successfully"})

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Route untuk membuat author
	http.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Handler untuk membuat author
			var req authorpb.CreateAuthorRequest

			// Parse request body sebagai JSON
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Panggil gRPC untuk membuat author
			resp, err := authorController.CreateAuthor(context.Background(), &req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Response sukses
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("HTTP server berjalan di port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Gagal menjalankan HTTP server: %v", err)
	}
}
