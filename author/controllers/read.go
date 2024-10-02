package controllers

import(
	"context"
	"sistem-microservice/author/models"
	authorpb "sistem-microservice/author/proto"

)

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