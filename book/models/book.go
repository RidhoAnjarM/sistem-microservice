package models

type Book struct {
	ID       int    `json:"id" gorm:"primaryKey"`              
	Title    string `json:"title"`
	Price    int    `json:"price"`
	AuthorID int    `json:"author_id"`                       
}
