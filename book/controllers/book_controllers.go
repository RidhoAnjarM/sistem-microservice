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

// NewBookController initializes a new BookController with a *gorm.DB instance
func NewBookController(db *gorm.DB) *BookController {
	return &BookController{
		DB: db,
	}
}

// CreateBook handles creating a new book in the database
func (bc *BookController) CreateBook(ctx context.Context, req *bookpb.CreateBookRequest) (*bookpb.CreateBookResponse, error) {
	// Create a new book instance based on the incoming request
	book := models.Book{
		Title:    req.GetTitle(),
		Price:    int(req.GetPrice()),
		AuthorID: int(req.GetAuthorId()),
	}

	// Insert the book into the database
	if err := bc.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	// Return the response with the created book details
	return &bookpb.CreateBookResponse{
		Book: &bookpb.Book{
			Id:       int32(book.ID),
			Title:    book.Title,
			Price:    int32(book.Price),
			AuthorId: int32(book.AuthorID),
		},
	}, nil
}

// GetBook handles fetching a book by its ID from the database
func (bc *BookController) GetBook(ctx context.Context, req *bookpb.GetBookRequest) (*bookpb.GetBookResponse, error) {
	var book models.Book

	// Fetch the book from the database using the ID from the request
	if err := bc.DB.First(&book, req.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	// Return the response with the fetched book details
	return &bookpb.GetBookResponse{
		Book: &bookpb.Book{
			Id:       int32(book.ID),
			Title:    book.Title,
			Price:    int32(book.Price),
			AuthorId: int32(book.AuthorID),
		},
	}, nil
}

// UpdateBook handles updating an existing book in the database
func (bc *BookController) UpdateBook(ctx context.Context, req *bookpb.UpdateBookRequest) (*bookpb.UpdateBookResponse, error) {
	var book models.Book

	// Check if the book exists in the database
	if err := bc.DB.First(&book, req.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	// Update the book fields with the data from the request
	book.Title = req.GetTitle()
	book.Price = int(req.GetPrice())

	// Save the updated book to the database
	if err := bc.DB.Save(&book).Error; err != nil {
		return nil, err
	}

	// Return the response with the updated book details
	return &bookpb.UpdateBookResponse{
		Book: &bookpb.Book{
			Id:       int32(book.ID),
			Title:    book.Title,
			Price:    int32(book.Price),
			AuthorId: int32(book.AuthorID),
		},
	}, nil
}

// DeleteBook handles removing a book from the database by its ID
func (bc *BookController) DeleteBook(ctx context.Context, req *bookpb.DeleteBookRequest) (*bookpb.DeleteBookResponse, error) {
	var book models.Book

	// Check if the book exists in the database
	if err := bc.DB.First(&book, req.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	// Delete the book from the database
	if err := bc.DB.Delete(&book).Error; err != nil {
		return nil, err
	}

	// Return a successful delete response
	return &bookpb.DeleteBookResponse{
		Message: "Book successfully deleted",
	}, nil
}
