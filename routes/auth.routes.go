package routes

import (
	"github.com/SSSBoOm/go-gorm-api/controller"
	"github.com/SSSBoOm/go-gorm-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/auth").Use(middleware.Auth)

	api.Get("/", controller.GetAllUsers)
	api.Get("/:id", controller.GetUsersById)

}
