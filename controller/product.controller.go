package controller

import (
	"github.com/SSSBoOm/go-gorm-api/database"
	"github.com/SSSBoOm/go-gorm-api/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	var Products []model.Product

	database.DB.Model(&model.Product{}).Find(&Products)
	return c.Status(200).JSON(&fiber.Map{
		"message": "Success",
		"success": true,
		"Data":    Products,
	})
}

func GetProductByID(c *fiber.Ctx) error {
	var Product model.Product
	id := c.Params("id", "0")

	if id == "0" {
		return c.Status(400).JSON(&fiber.Map{
			"message": "No ID",
			"success": false,
		})
	}

	if res := database.DB.Model(&model.Product{}).Where("id = ?", id).First(&Product); res.RowsAffected == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Not Found",
			"success": false,
			"Data":    nil,
		})
	} else {
		return c.Status(200).JSON(&fiber.Map{
			"message": "Success",
			"success": true,
			"Data":    Product,
		})
	}
}

func CreateProduct(c *fiber.Ctx) error {
	Product := new(model.Product)

	if err := c.BodyParser(Product); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	if res := database.DB.Model(&model.Product{}).Create(Product); res.Error != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": res.Error.Error(),
			"success": false,
		})
	}

	return c.Status(201).JSON(&fiber.Map{
		"message": "Created",
		"success": true,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	Product := new(model.Product)
	id := c.Params("id", "0")

	if id == "0" {
		return c.Status(400).JSON(&fiber.Map{
			"message": "No ID",
			"success": false,
		})
	}

	if err := c.BodyParser(Product); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	if res := database.DB.Model(&model.Product{}).Where("id= ?", id).Updates(Product); res.Error != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": res.Error.Error(),
			"success": false,
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"message": "Updated",
		"success": true,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	Product := new(model.Product)
	id := c.Params("id", "0")

	if id == "0" {
		return c.Status(400).JSON(&fiber.Map{
			"message": "No ID",
			"success": false,
		})
	}

	if res := database.DB.Model(&model.Product{}).Where("id= ?", id).Delete(&Product); res.Error != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": res.Error.Error(),
			"success": false,
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"message": "Deleted",
		"success": true,
	})
}
