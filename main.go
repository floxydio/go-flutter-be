package main

import (
	controller "gohtml/Controllers"
	database "gohtml/Database"
	"log"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	app := fiber.New()

	database.Connect()

	app.Get("/", controller.Homepage)

	// Login
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	// Task
	app.Get("/api/task", controller.Task)
	app.Post("/api/create/task", controller.PostTodo)
	app.Put("/api/task/:id", controller.PutTask)
	app.Delete("/api/deletetask/:id", controller.DeleteTask)
	log.Fatal(app.Listen(":2000"))

}
