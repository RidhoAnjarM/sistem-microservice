package main

import (
    "fmt"
    "log"
    "net/http"
    authorController "sistem-microservice/author/controllers"
    bookController "sistem-microservice/book/controllers"
)

func main() {
    fmt.Println("Gateway Starting...")

    http.HandleFunc("/authors", authorHandler)
    http.HandleFunc("/books", bookHandler)

    fmt.Println("Gateway is running on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}

// Handler untuk author, memanggil controller dari author/controllers
func authorHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        authorController.CreateAuthor(w, r)
    case http.MethodGet:
        idStr := r.URL.Query().Get("id")
        if idStr != "" {
            authorController.GetAuthor(w, r)
        } else {
            authorController.GetAllAuthors(w, r)
        }
    case http.MethodPut:
        authorController.UpdateAuthor(w, r)
    case http.MethodDelete:
        authorController.DeleteAuthor(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// Handler untuk book, memanggil controller dari book/controllers
func bookHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        bookController.CreateBook(w, r)
    case http.MethodGet:
        idStr := r.URL.Query().Get("id")
        if idStr != "" {
            bookController.GetBook(w, r)
        } else {
            bookController.GetAllBooks(w, r)
        }
    case http.MethodPut:
        bookController.UpdateBook(w, r)
    case http.MethodDelete:
        bookController.DeleteBook(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
