package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(app *fiber.App) {
	route := a.Group("/swagger")

	route.Get("*", swagger.Handler)
}