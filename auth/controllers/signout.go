package controllers

import "github.com/gofiber/fiber"

// Signout logs the user out
func Signout(c *fiber.Ctx) {
	c.Send("Signing user out")
}
