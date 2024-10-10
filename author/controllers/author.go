package controllers

import (
	"gorm.io/gorm"
	authorpb "sistem-microservice/author/proto"
)

type AuthorController struct {
	authorpb.UnimplementedAuthorServiceServer
	DB *gorm.DB
}

func NewAuthorController(db *gorm.DB) *AuthorController {
	return &AuthorController{DB: db}
}