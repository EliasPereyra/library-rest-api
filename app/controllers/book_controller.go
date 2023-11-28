package controllers

import (
	"library-rest-api/app/models"
	"library-rest-api/pkg/utils"
	"time"

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

	// Create a book
	// @Description Create a new book
	// @Summary create a new book
	// @Tags Book
	// @Accept json
	// @Produce json
	// @Param title body string true "Title"
	// @Param author body string true "Author"
	// @Param book_attrs body models.BookAttrs true "Book attributes"
	// @Sucess 200 {object} models.Book
	// @Security ApiKeyAuth
	// @Router /v1/book [post]

	func CreateBook(c *fiber.Ctx) error {
		now := time.Now().Unix()

		claims, err := utils.ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.App{
				"error": true,
				"msg": err.Error(),
			})
		}

		expires := claims.Expires
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg": "Unauthorized, the token has expired",
		})
	}

	book := &models.Book{}

	if err := c.BodyParaser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": err.Error()
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg": err.Error()
		})
	}

	validate := utils.NewValidator()

	book.ID = uuid.New()
	book.CreatedAt = time.Now()
	book.BookStatus = 1

	if err := validate.Struct(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": utils.ValidatorErrors(err)
		})
	}

	if err := db.CreateBook(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg": err.Error()
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg": nil,
		"book": book
	})
	}

	// Update a book
	// @Description Update a book
	// @Summary update a book
	// @Tags Book
	// @Accept json
	// @Produce json
	// @Param id body string true "Book ID"
	// @Param title body string true "Title"
	// @Param author body string true "Author"
	// @Param book_status body integer true "Book status"
	// @Param book_attrs body models.BookAttrs true "Book attributes"
	// @Success 201 {string} status "ok"
	// @Security ApiKeyAuth
	// @Router /v1/book [put]
	func UpdateBook(c *fiber.Ctx) err {
		now := time.Now().Unix()

		claims, err := utils.ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"err": true,
				"msg": err.Error()
			})
		}

		expires := claims.Expires

		if now > expires {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg": "Unauthorized: check the expiration time of your token",
			})
		}

		book := &models.Book

		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

		bookFound, err := db.GetBook(book.ID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg": "The book with given ID doesn't exist",
			})
		}

		book.UpdatedAt = time.Now()

		validate := utils.NewValidator()

		if err := validate.Struct(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg": utils.ValidatorErrors(err),
			})
		}

		if err := db.UpdateBook(bookFound.ID, book); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}

	// Delete a book
	// @Description Delete a book with a given ID
	// @Summary delete a book with a given ID
	// @Tags Book
	// @Accept json
	// @Produce json
	// @Param id body string true "Book ID"
	// @Sucess 204 {string} status "ok"
	// @Security ApiKeyAuth
	// @Router /v1/book [delete]
	func DeleteBook(c *fiber.Ctx) error {
		now := time.Now().Unix()

		claims, err := utils.ExtractTokenMetada(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg": err.Error(),
			})
		}

		expires := claims.Expires

		if now > expires {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg": "Unauthorized: check the expiration time of your token",
			})
		}

		book := &models.Book{}

		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg": err.Error(),
			})
		}

		validate := utils.NewValidator()

		if err := validate.StructPartial(book, "id"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg": utils.ValidatorErrors(err),
			})
		}

		db, err := database.OpenDBConnection()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg": err.Error(),
			})
		}

		bookFound, err := db.GetBook(book.ID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg": "The book with the given ID wasn't found",
			})
		}

		if err := db.DeleteBook(bookFound.ID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}