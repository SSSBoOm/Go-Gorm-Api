package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	c.SendStatus(200)
	return nil
}
