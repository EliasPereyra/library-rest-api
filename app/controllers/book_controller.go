package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Get all books
// @Description Get all the books
// @Summary get all the books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]

	func GetBooks(c *fiber.Ctx) error {
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

	// Get a book
	// @Description Get a book by an id
	// @Summary get a book
	// @Tags Book
	// @Accept json
	// @Produce json
	// @Param id path string true "Book ID"
	// @Success 200 {object} models.Book
	// @Router /v1/book/{id} [get]
	func GetBook(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg": err.Error(),
			})
		}

		db, err := database.OpenDBConnection()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg": err.Error(),
			})
		}

		book, err := db.GetBook(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg": "The book wasn't found",
				"book": nil,
			})
		}

		return c.Status(fiber.Map{
			"error": false,
			"msg": nil,
			"book": book,
		})
	}