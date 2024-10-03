package controllers

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"sistem-microservice/author/models"
	authorpb "sistem-microservice/author/proto"
	"strconv"
	"gorm.io/gorm"
)

type AuthorController struct {
	authorpb.UnimplementedAuthorServiceServer
	DB *gorm.DB
}

func NewAuthorController(db *gorm.DB) *AuthorController {
	return &AuthorController{DB: db}
}


// GetAuthorHandler
func (ac *AuthorController) GetAuthorHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
    return
}
req := &authorpb.GetAuthorRequest{Id: int32(id)}

	res, err := ac.GetAuthorWithBooks(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Delete author and cascade delete books
func (ac *AuthorController) DeleteAuthorHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var author models.Author
	if err := ac.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	ac.DB.Delete(&author)
	c.JSON(http.StatusOK, gin.H{"message": "Author and related books deleted successfully"})
}

