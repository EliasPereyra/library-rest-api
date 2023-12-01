package routes

import (
	"library-rest-api/app/controllers"
	"library-rest-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook())

	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook())

	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook())
}