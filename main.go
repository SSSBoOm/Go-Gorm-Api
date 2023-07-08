package main

import (
	"github.com/SSSBoOm/go-gorm-api/database"
	"github.com/SSSBoOm/go-gorm-api/router"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID    uint
	Code  string
	Price uint
}

func main() {
	if err := database.Connect(); err != nil {
		panic("Cant Connect Database")
	}

	app := fiber.New()

	router.MainRoutes(app)

	app.Listen(":3000")
}
