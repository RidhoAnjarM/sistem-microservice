package controllers

import (
	"context"
	"errors"
	"fmt"
	"sistem-microservice/book/models"
	bookpb "sistem-microservice/book/proto"

	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
	bookpb.UnimplementedBookServiceServer
}

// NewBookController 
func NewBookController(db *gorm.DB) *BookController {
	return &BookController{
		DB: db,
	}
}

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
