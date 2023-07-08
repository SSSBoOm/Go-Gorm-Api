package controller

import (
	"github.com/SSSBoOm/go-gorm-api/database"
	"github.com/SSSBoOm/go-gorm-api/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	var Users []model.User

	database.DB.Model(&model.User{}).Find(&Users)
	return c.Status(200).JSON(&fiber.Map{
		"message": "Success",
		"success": true,
		"Data":    Users,
		"ID":      c.Locals("USER"),
	})
}

func GetUsersById(c *fiber.Ctx) error {
	var Users []model.User
	id := c.Params("id", "0")

	if id == "0" {
		return c.Status(400).JSON(&fiber.Map{
			"message": "No ID",
			"success": false,
		})
	}

	database.DB.Model(&model.User{}).Where("id = ", id).Find(&Users)
	return c.Status(200).JSON(&fiber.Map{
		"message": "Success",
		"success": true,
		"Data":    Users,
	})
}
