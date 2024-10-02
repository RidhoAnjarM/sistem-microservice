package controllers

import (
	"context"
	"errors"
	"fmt"
	"sistem-microservice/book/models"
	bookpb "sistem-microservice/book/proto"

	"gorm.io/gorm"
)

// UpdateBook
func (bc *BookController) UpdateBook(ctx context.Context, req *bookpb.UpdateBookRequest) (*bookpb.UpdateBookResponse, error) {
	var book models.Book

	// Check data di database
	if err := bc.DB.First(&book, req.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	// Update book 
	book.Title = req.GetTitle()
	book.Price = int(req.GetPrice())

	// Simpan update book ke database
	if err := bc.DB.Save(&book).Error; err != nil {
		return nil, err
	}

	// Return response
	return &bookpb.UpdateBookResponse{
		Book: &bookpb.Book{
			Id:       int32(book.ID),
			Title:    book.Title,
			Price:    int32(book.Price),
			AuthorId: int32(book.AuthorID),
		},
	}, nil
}