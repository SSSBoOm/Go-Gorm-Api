package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	h := c.Get("Authorization")

	if h == "" {
		return c.Status(401).JSON(&fiber.Map{
			"message": "Unauthorized",
			"success": false,
		})
	}

	return c.Next()
}
