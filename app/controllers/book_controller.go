package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Get all books
// @Description Get all the books
// @Summary get all the books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]

	func getBooks(c *fiber.Ctx) error {
		db, err := database.OpenDBConnection()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg": err.Error(),
			})
		}

		books, err := db.GetBooks()

		if err != nil {
			c.Status(fiber.StatusNotFound).JSON(fiber.App{
				"error": true,
				"msg": "There are no books",
				"count": 0,
				"books": nil,
			})
		}

		return c.JSON(fiber.Map{
			"error": false,
			"msg": nil,
			"count": len(books),
			"books": books,
		})
	}