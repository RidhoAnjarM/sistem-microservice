package controllers

import(
	"context"
	"sistem-microservice/author/models"
	authorpb "sistem-microservice/author/proto"

)


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