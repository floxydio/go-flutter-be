package model

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
