package controllers

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
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

