package models

type Book struct {
	ID       int    `gorm:"primaryKey"`
	Title    string `gorm:"type:varchar(255)"`
	Price    int
	AuthorID int
}
