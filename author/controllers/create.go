package controllers

import(
	"context"
	"sistem-microservice/author/models"
	authorpb "sistem-microservice/author/proto" 
)



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