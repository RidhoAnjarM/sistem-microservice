package controllers

import (
	"context"
	"github.com/raffa/book/models"
	bookpb "github.com/raffa/book/proto"
)

// CreateBook
func (bc *BookController) CreateBook(ctx context.Context, req *bookpb.CreateBookRequest) (*bookpb.CreateBookResponse, error) {
	// Create a new book instance based on the incoming request
	book := models.Book{
		Title:    req.GetTitle(),
		Price:    int(req.GetPrice()),
		AuthorID: int(req.GetAuthorId()),
	}

	// Insert buku ke database
	if err := bc.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	// Return response
	return &bookpb.CreateBookResponse{
		Book: &bookpb.Book{
			Id:       int32(book.ID),
			Title:    book.Title,
			Price:    int32(book.Price),
			AuthorId: int32(book.AuthorID),
		},
	}, nil
}