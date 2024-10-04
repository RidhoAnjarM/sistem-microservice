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

// CreateAuthorHandler
func (ac *AuthorController) CreateAuthorHandler(c *gin.Context) {
	var req authorpb.CreateAuthorRequest

	// Parse request dari body JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Panggil fungsi gRPC untuk membuat author
	resp, err := ac.CreateAuthor(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kirim response sebagai JSON
	c.JSON(http.StatusOK, resp)
}

// Delete author and cascade delete books
func (ac *AuthorController) DeleteAuthorHandler(c *gin.Context) {
    // Mengambil ID author dari parameter URL dan mengonversinya ke int
    id, _ := strconv.Atoi(c.Param("id"))
    
    var author models.Author
    // Mencari author berdasarkan ID
    if err := ac.DB.First(&author, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
        return
    }

    // Menghapus semua buku yang terkait dengan author sebelum menghapus author
    if err := ac.DB.Where("author_id = ?", id).Delete(&models.Book{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete related books"})
        return
    }
    
    // Menghapus author setelah buku-buku terkait dihapus
    if err := ac.DB.Delete(&author).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author"})
        return
    }

    // Merespons sukses
    c.JSON(http.StatusOK, gin.H{"message": "Author and related books deleted successfully"})
}


