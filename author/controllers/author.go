package controllers

import (
	"context"
	"sistem-microservice/author/models"
	authorpb "sistem-microservice/author/proto"

	"gorm.io/gorm"
)

type AuthorController struct {
	authorpb.UnimplementedAuthorServiceServer
	DB *gorm.DB
}

func NewAuthorController(db *gorm.DB) *AuthorController {
	return &AuthorController{DB: db}
}

// Create Author
func (c *AuthorController) CreateAuthor(ctx context.Context, req *authorpb.CreateAuthorRequest) (*authorpb.CreateAuthorResponse, error) {
	author := models.Author{Name: req.Name, Email: req.Email}
	result := c.DB.Create(&author)

	if result.Error != nil {
		return nil, result.Error
	}

	return &authorpb.CreateAuthorResponse{
		Status:  "success",
		Message: "Author created successfully",
		Author: &authorpb.Author{
			Id:    int32(author.ID),
			Name:  author.Name,
			Email: author.Email,
		},
	}, nil
}

// Get Author
func (c *AuthorController) GetAuthor(ctx context.Context, req *authorpb.GetAuthorRequest) (*authorpb.GetAuthorResponse, error) {
	var author models.Author
	if err := c.DB.First(&author, req.Id).Error; err != nil {
		return nil, err
	}

	return &authorpb.GetAuthorResponse{
		Status:  "success",
		Message: "Detail author berhasil ditemukan",
		Author: &authorpb.Author{
			Id:    int32(author.ID),
			Name:  author.Name,
			Email: author.Email,
		},
	}, nil
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

// Update Author
func (c *AuthorController) UpdateAuthor(ctx context.Context, req *authorpb.UpdateAuthorRequest) (*authorpb.UpdateAuthorResponse, error) {
	var author models.Author
	if err := c.DB.First(&author, req.Id).Error; err != nil {
		return nil, err
	}

	author.Name = req.Name
	author.Email = req.Email
	c.DB.Save(&author)

	return &authorpb.UpdateAuthorResponse{
		Status:  "success",
		Message: "Author updated successfully",
		Author: &authorpb.Author{
			Id:    req.Id,
			Name:  author.Name,
			Email: author.Email,
		},
	}, nil
}

// Delete Author
func (c *AuthorController) DeleteAuthor(ctx context.Context, req *authorpb.DeleteAuthorRequest) (*authorpb.DeleteAuthorResponse, error) {
	if err := c.DB.Delete(&models.Author{}, req.Id).Error; err != nil {
		return nil, err
	}

	return &authorpb.DeleteAuthorResponse{
		Status:  "success",
		Message: "Author deleted successfully",
	}, nil
}
