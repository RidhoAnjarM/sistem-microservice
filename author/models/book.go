package models

type Book struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Title    string `gorm:"not null" json:"title"`
    Price    uint   `json:"price"`
    AuthorID uint   `gorm:"not null" json:"author_id"`
}
