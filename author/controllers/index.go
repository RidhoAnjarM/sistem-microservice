package controllers

import (
	"sistem-microservice/author/models"

	"gorm.io/gorm"
)

type InDb struct {
	sql *gorm.DB
}

func NewInstance() InDb {
	db, err := models.GetSqlConnection()
	if err != nil {
		panic(err)
	}

	return InDb{
		sql: db,
	}
}

// Ping : pinging db
func (idb InDb) Ping() error {
	db, err := idb.sql.DB()
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}
