package models

type User struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt string `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt string `json:"deletedAt" gorm:"column:deletedAt"`
}

func (User) TableName() string {
	return "users"
}
