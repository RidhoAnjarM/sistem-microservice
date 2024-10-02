package controllers

import(
	"context"
	"sistem-microservice/author/models"
	authorpb "sistem-microservice/author/proto"

)

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