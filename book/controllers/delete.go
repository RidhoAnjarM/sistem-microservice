package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/raffa/book/models"
	bookpb "github.com/raffa/book/proto"

	"gorm.io/gorm"
)

// DeleteBook
func (bc *BookController) DeleteBook(ctx context.Context, req *bookpb.DeleteBookRequest) (*bookpb.DeleteBookResponse, error) {
	var book models.Book

	// Check data di database
	if err := bc.DB.First(&book, req.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	// Delete book dari database
	if err := bc.DB.Delete(&book).Error; err != nil {
		return nil, err
	}

	// Return response
	return &bookpb.DeleteBookResponse{
		Message: "Book successfully deleted",
	}, nil
}