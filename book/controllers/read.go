package controllers

import (
	"context"
	"errors"
	"fmt"
	"sistem-microservice/book/models"
	bookpb "sistem-microservice/book/proto"

	"gorm.io/gorm"
)

// GetBook
func (bc *BookController) GetBook(ctx context.Context, req *bookpb.GetBookRequest) (*bookpb.GetBookResponse, error) {
	var book models.Book

	// Mengambil data buku di database
	if err := bc.DB.First(&book, req.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	// Return response
	return &bookpb.GetBookResponse{
		Book: &bookpb.Book{
			Id:       int32(book.ID),
			Title:    book.Title,
			Price:    int32(book.Price),
			AuthorId: int32(book.AuthorID),
		},
	}, nil
}

// GetBooksByAuthorId
func (bc *BookController) GetBooksByAuthorId(ctx context.Context, req *bookpb.GetBooksByAuthorIdRequest) (*bookpb.GetBooksByAuthorIdResponse, error) {
	var books []models.Book
	if err := bc.DB.Where("author_id = ?", req.GetAuthorId()).Find(&books).Error; err != nil {
		return nil, err
	}

	var bookList []*bookpb.Book
	for _, book := range books {
		bookList = append(bookList, &bookpb.Book{
			Id:       int32(book.ID),
			Title:    book.Title,
			Price:    int32(book.Price),
			AuthorId: int32(book.AuthorID),
		})
	}

	return &bookpb.GetBooksByAuthorIdResponse{
		Books: bookList,
	}, nil
}