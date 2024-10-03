package models

type Author struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"unique,not null" json:"email"`
	Books []Book `gorm:"foreignKey:AuthorID"`
}


