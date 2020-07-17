package controllers

import "github.com/gofiber/fiber"

// Ping tests the API
func Ping(c *fiber.Ctx) {
	c.Send("Pong!")
}
