package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// method for generating a new access token
// @Description create a new access token
// @Summary create a new access token
// @Tags token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/token/new [get]
func GetNewAccessToken(c *fiber.Ctx) error {
	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg": nil,
		"token": token,
	})
}