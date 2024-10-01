package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetSqlConnection() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal Menyambung Ke Database")
	}

	return db, nil
}
