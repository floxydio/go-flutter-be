package database

import (
	"fmt"

	model "gohtml/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost)/lmsgo"), &gorm.Config{})

	if err != nil {
		panic("This DB Not Connected")
	} else {
		fmt.Println("Connected")
	}

	DB = db

	db.AutoMigrate(model.Task{})
	db.AutoMigrate(model.User{})

}
