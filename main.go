package main

import "github.com/gofiber/fiber/v2"

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService https://swagger.io/terms
// @contact.name API Support
// @contact.email mail@mail.com
// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.SwaggerRoutes(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServerWithGraceShutdown(app)
}