package model

type Task struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
}
