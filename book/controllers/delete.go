package controllers

import (
	"context"
	"errors"
	"fmt"
	"sistem-microservice/book/models"
	bookpb "sistem-microservice/book/proto"

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
		Status: "success",
		Message: "Book successfully deleted",
	}, nil
}

// Di server BookService
func (bc *BookController) DeleteBooksByAuthorId(ctx context.Context, req *bookpb.DeleteBooksByAuthorIdRequest) (*bookpb.DeleteBooksByAuthorIdResponse, error) {
    authorId := req.GetAuthorId()

    // Menghapus semua buku berdasarkan author_id
    if err := bc.DB.Where("author_id = ?", authorId).Delete(&models.Book{}).Error; err != nil {
        return &bookpb.DeleteBooksByAuthorIdResponse{
            Status:  "error",
            Message: "Failed to delete related books",
        }, err
    }

    return &bookpb.DeleteBooksByAuthorIdResponse{
        Status:  "success",
        Message: "Related books deleted successfully",
    }, nil
}
