package controllers

import (
	bookpb "sistem-microservice/book/proto"

	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
	bookpb.UnimplementedBookServiceServer
}

// NewBookController 
func NewBookController(db *gorm.DB) *BookController {
	return &BookController{
		DB: db,
	}
}
