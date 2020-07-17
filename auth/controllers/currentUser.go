package controllers

import (
	"github.com/gofiber/fiber"
)

// CurrentUser get a specific user
func CurrentUser(c *fiber.Ctx) {
	id := c.Params("id")

	_ = c.JSON(&fiber.Map{
		"success": true,
		"data":    id,
	})
}
