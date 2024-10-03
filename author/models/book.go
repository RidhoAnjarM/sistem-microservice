package models

type Book struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Title    string `gorm:"not null" json:"title"`
    Price    int    `json:"price"`
    AuthorID uint   `json:"author_id"`
}
