package controller

import (
	database "gohtml/Database"
	model "gohtml/Models"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.SendString("Its Me Dio")

}

func PostTodo(c *fiber.Ctx) error {
	quiz := &model.Task{
		Task: c.FormValue("task"),
	}

	res := &model.Result{
		Code:    200,
		Data:    quiz,
		Message: "Successfully Add Task",
	}

	database.DB.Create(&quiz)

	return c.JSON(res)

}

func Task(c *fiber.Ctx) error {
	gettask := []model.Task{}

	database.DB.Find(&gettask)

	return c.JSON(fiber.Map{
		"data":    gettask,
		"message": "List Todo",
	})
}

func PutTask(c *fiber.Ctx) error {
	id := c.Params("id")

	task := model.Task{
		Task: c.FormValue("task"),
	}

	database.DB.Model(&task).Where("id = ?", id).Updates(&task)

	return c.JSON(fiber.Map{
		"message": "Success Changed This Post",
	})

}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	var task model.Task

	database.DB.First(&task, id)
	database.DB.Delete(&task)

	return c.JSON(fiber.Map{
		"message": "Berhasil Delete",
	})
}

// User

func Register(c *fiber.Ctx) error {
	user := model.User{
		Email:    c.FormValue("email"),
		Name:     c.FormValue("name"),
		Password: c.FormValue("password"),
	}
	database.DB.Create(&user)

	if user.Name == c.FormValue("name") && user.Email == c.FormValue("email") && user.Password == c.FormValue("password") {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":   200,
			"name":     user.Name,
			"email":    user.Email,
			"password": user.Password,
			"message":  "Register Successfully",
		})
	}

	c.Type("json", "utf-8")

	return c.JSON(&user)

}

func Login(c *fiber.Ctx) error {
	var user model.User

	database.DB.Where("email = ? AND password = ?", c.FormValue("email"), c.FormValue("password")).First(&user)

	if user.Email == c.FormValue("email") && user.Password == c.FormValue("password") {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":   200,
			"email":    user.Email,
			"password": user.Password,
			"name":     user.Name,
		})
	}

	if user.Email == c.FormValue("") || user.Password == c.FormValue("") {
		return c.JSON(fiber.Map{
			"status":  401,
			"message": "Kosong ??",
		})
	}

	if user.Email != c.FormValue("email") && user.Password != c.FormValue("password") {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Login Failed",
		})
	}

	c.Accepts("application/json")
	return c.JSON(user)
}
