package main

import (
	"fmt"

	"github.com/SSSBoOm/go-gorm-api/database"
	"github.com/SSSBoOm/go-gorm-api/routes"
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
		fmt.Println("Cant Connect Database.")
		panic(err)
	}

	app := fiber.New()

	routes.AuthRoutes(app)
	routes.ProductRoutes(app)

	app.Listen(":3000")
}
