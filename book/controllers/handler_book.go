package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "strconv"
    "google.golang.org/grpc"
    bookpb "sistem-microservice/book/proto"
)

const bookServiceURL = "localhost:50052"

// CreateBook: Membuat buku baru
func CreateBook(w http.ResponseWriter, r *http.Request) {
    // Dekode body request menjadi map
    var book map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    conn, err := grpc.Dial(bookServiceURL, grpc.WithInsecure())
    if err != nil {
        http.Error(w, "Error connecting to book service", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    client := bookpb.NewBookServiceClient(conn)

    req := &bookpb.CreateBookRequest{
        Title:    book["title"].(string),
        Price:    int32(book["price"].(float64)),
        AuthorId: int32(book["author_id"].(float64)),
    }

    resp, err := client.CreateBook(context.Background(), req)
    if err != nil {
        http.Error(w, "Error creating book", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":  "success",
        "message": "Book created successfully",
        "book": map[string]interface{}{
            "id":        resp.Book.Id,
            "title":     resp.Book.Title,
            "price":     resp.Book.Price,
            "author_id": resp.Book.AuthorId,
        },
    })
}

// GetBook: Mendapatkan buku berdasarkan ID
func GetBook(w http.ResponseWriter, r *http.Request) {
    conn, err := grpc.Dial(bookServiceURL, grpc.WithInsecure())
    if err != nil {
        http.Error(w, "Error connecting to book service", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    client := bookpb.NewBookServiceClient(conn)

    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    req := &bookpb.GetBookRequest{Id: int32(id)}
    resp, err := client.GetBook(context.Background(), req)
    if err != nil {
        http.Error(w, "Error getting book", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":  "success",
        "message": "Book retrieved successfully",
        "book": map[string]interface{}{
            "id":        resp.Book.Id,
            "title":     resp.Book.Title,
            "price":     resp.Book.Price,
            "author_id": resp.Book.AuthorId,
        },
    })
}

// GetAllBooks: Mendapatkan semua buku
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
    conn, err := grpc.Dial(bookServiceURL, grpc.WithInsecure())
    if err != nil {
        http.Error(w, "Error connecting to book service", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    client := bookpb.NewBookServiceClient(conn)

    req := &bookpb.Empty{} // Request untuk mendapatkan semua buku
    resp, err := client.GetAllBooks(context.Background(), req)
    if err != nil {
        http.Error(w, "Error getting books", http.StatusInternalServerError)
        return
    }

    books := make([]map[string]interface{}, len(resp.Book))
    for i, book := range resp.Book {
        books[i] = map[string]interface{}{
            "id":        book.Id,
            "title":     book.Title,
            "price":     book.Price,
            "author_id": book.AuthorId,
        }
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":  "success",
        "message": "Books retrieved successfully",
        "books":   books,
    })
}

// UpdateBook: Mengupdate buku berdasarkan ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    // Dekode body request menjadi map
    var book map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    conn, err := grpc.Dial(bookServiceURL, grpc.WithInsecure())
    if err != nil {
        http.Error(w, "Error connecting to book service", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    client := bookpb.NewBookServiceClient(conn)

    id, err := strconv.Atoi(book["id"].(string)) // Mengambil ID dari body
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    req := &bookpb.UpdateBookRequest{
        Id:    int32(id),
        Title: book["title"].(string),
        Price: int32(book["price"].(float64)),
    }

    resp, err := client.UpdateBook(context.Background(), req)
    if err != nil {
        http.Error(w, "Error updating book", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":  "success",
        "message": "Book updated successfully",
        "book": map[string]interface{}{
            "id":        resp.Book.Id,
            "title":     resp.Book.Title,
            "price":     resp.Book.Price,
            "author_id": resp.Book.AuthorId,
        },
    })
}

// DeleteBook: Menghapus buku berdasarkan ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    conn, err := grpc.Dial(bookServiceURL, grpc.WithInsecure())
    if err != nil {
        http.Error(w, "Error connecting to book service", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    client := bookpb.NewBookServiceClient(conn)

    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    req := &bookpb.DeleteBookRequest{Id: int32(id)}
    _, err = client.DeleteBook(context.Background(), req)
    if err != nil {
        http.Error(w, "Error deleting book", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":  "success",
        "message": "Book deleted successfully",
    })
}
