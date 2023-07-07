package router

import (
	"github.com/SSSBoOm/go-gorm-api/controller"
	"github.com/gofiber/fiber/v2"
)

func MainRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", controller.GetAllProducts)
	api.Get("/:id", controller.GetProductByID)
	api.Post("/", controller.CreateProduct)
	api.Put("/:id", controller.UpdateProduct)
	api.Delete("/:id", controller.DeleteProduct)
}
