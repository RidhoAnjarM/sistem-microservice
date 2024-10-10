package controllers

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"net/http"
	authorpb "sistem-microservice/author/proto"
	bookpb "sistem-microservice/book/proto"
	"strconv"
)

const authorServiceURL = "localhost:50051"
const bookServiceURL = "localhost:50052"

// CreateAuthor: Membuat author baru
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(authorServiceURL, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Error connecting to author service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := authorpb.NewAuthorServiceClient(conn)

	req := &authorpb.CreateAuthorRequest{
		Name:  author["name"].(string),
		Email: author["email"].(string),
	}

	resp, err := client.CreateAuthor(context.Background(), req)
	if err != nil {
		http.Error(w, "Error creating author", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Author created successfully",
		"author": map[string]interface{}{
			"id":    resp.Author.Id,
			"name":  resp.Author.Name,
			"email": resp.Author.Email,
		},
	})
}

// GetAuthor: Mendapatkan author berdasarkan ID
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(authorServiceURL, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Error connecting to author service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := authorpb.NewAuthorServiceClient(conn)

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	req := &authorpb.GetAuthorRequest{Id: int32(id)}
	resp, err := client.GetAuthor(context.Background(), req)
	if err != nil {
		http.Error(w, "Error getting author", http.StatusInternalServerError)
		return
	}

	// Mengambil buku berdasarkan author_id
	bookConn, err := grpc.Dial(bookServiceURL, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Error connecting to book service", http.StatusInternalServerError)
		return
	}
	defer bookConn.Close()

	bookClient := bookpb.NewBookServiceClient(bookConn)

	bookReq := &bookpb.GetBooksByAuthorIdRequest{AuthorId: int32(id)}
	bookResp, err := bookClient.GetBooksByAuthorId(context.Background(), bookReq)
	if err != nil {
		http.Error(w, "Error getting books", http.StatusInternalServerError)
		return
	}

	// Menyusun respons JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Author retrieved successfully",
		"author": map[string]interface{}{
			"id":    resp.Author.Id,
			"name":  resp.Author.Name,
			"email": resp.Author.Email,
		},
		"books": bookResp.Books,
	})
}

// GetAllAuthors: Mendapatkan semua author
func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(authorServiceURL, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Error connecting to author service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := authorpb.NewAuthorServiceClient(conn)

	req := &authorpb.Empty{} // Request kosong untuk mendapatkan semua penulis
	resp, err := client.GetAllAuthors(context.Background(), req)
	if err != nil {
		http.Error(w, "Error getting authors", http.StatusInternalServerError)
		return
	}

	authors := make([]map[string]interface{}, len(resp.Authors))
	for i, author := range resp.Authors {
		authors[i] = map[string]interface{}{
			"id":    author.Id,
			"name":  author.Name,
			"email": author.Email,
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Authors retrieved successfully",
		"authors": authors,
	})
}

// UpdateAuthor: Mengupdate author berdasarkan ID
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var author map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(authorServiceURL, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Error connecting to author service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := authorpb.NewAuthorServiceClient(conn)

	
	idFloat, ok := author["id"].(float64) 
	if !ok {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	req := &authorpb.UpdateAuthorRequest{
		Id:    int32(idFloat), 
		Name:  author["name"].(string),
		Email: author["email"].(string),
	}

	resp, err := client.UpdateAuthor(context.Background(), req)
	if err != nil {
		http.Error(w, "Error updating author", http.StatusInternalServerError)
		return
	}

	// Membangun respons
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Author updated successfully",
		"author": map[string]interface{}{
			"id":    resp.Author.Id,
			"name":  resp.Author.Name,
			"email": resp.Author.Email,
		},
	})
}

// DeleteAuthor: Menghapus author dan buku terkait berdasarkan ID
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(authorServiceURL, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Error connecting to author service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := authorpb.NewAuthorServiceClient(conn)

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	// Menghapus buku terkait terlebih dahulu
	bookConn, err := grpc.Dial(bookServiceURL, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Error connecting to book service", http.StatusInternalServerError)
		return
	}
	defer bookConn.Close()

	bookClient := bookpb.NewBookServiceClient(bookConn)
	bookReq := &bookpb.DeleteBooksByAuthorIdRequest{AuthorId: int32(id)}
	_, err = bookClient.DeleteBooksByAuthorId(context.Background(), bookReq)
	if err != nil {
		http.Error(w, "Error deleting related books", http.StatusInternalServerError)
		return
	}

	// Menghapus author setelah buku dihapus
	req := &authorpb.DeleteAuthorRequest{Id: int32(id)}
	_, err = client.DeleteAuthor(context.Background(), req)
	if err != nil {
		http.Error(w, "Error deleting author", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Author and related books deleted successfully",
	})
}
