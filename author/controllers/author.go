package controllers

import (
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

