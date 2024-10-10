package controllers

import (
	"context"
	"errors"
	"fmt"
	"sistem-microservice/author/models"
	authorpb "sistem-microservice/author/proto"
	bookpb "sistem-microservice/book/proto"

	"gorm.io/gorm"
)

// GetAuthor
func (ac *AuthorController) GetAuthor(ctx context.Context, req *authorpb.GetAuthorRequest) (*authorpb.GetAuthorResponse, error) {
	var author models.Author
	// Ambil author dari database
	if err := ac.DB.Preload("Books").First(&author, req.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("author not found")
		}
		return nil, err
	}

	// Buat response
	return &authorpb.GetAuthorResponse{
		Status:  "success",
		Message: "Detail author berhasil ditemukan",
		Author: &authorpb.Author{
			Id:    int32(author.ID),
			Name:  author.Name,
			Email: author.Email,
			Books: convertBooks(author.Books),
		},
	}, nil
}

// Fungsi untuk mengonversi daftar buku ke format proto
func convertBooks(books []models.Book) []*bookpb.Book {
	var bookList []*bookpb.Book
	for _, b := range books {
		bookList = append(bookList, &bookpb.Book{
			Id:       int32(b.ID),
			Title:    b.Title,
			Price:    int32(b.Price),
			AuthorId: int32(b.AuthorID),
		})
	}
	return bookList
}

// Get All Authors
func (c *AuthorController) GetAllAuthors(ctx context.Context, req *authorpb.Empty) (*authorpb.GetAllAuthorsResponse, error) {
	var authors []models.Author
	result := c.DB.Find(&authors)

	if result.Error != nil {
		return nil, result.Error
	}

	var authorList []*authorpb.Author
	for _, author := range authors {
		authorList = append(authorList, &authorpb.Author{
			Id:    int32(author.ID),
			Name:  author.Name,
			Email: author.Email,
		})
	}

	return &authorpb.GetAllAuthorsResponse{
		Status:  "success",
		Message: "Authors retrieved successfully",
		Authors: authorList,
	}, nil
}

