package controllers

import "github.com/gofiber/fiber"

// Signin logs the user in
func Signin(c *fiber.Ctx) {
	c.Send("Signing user in")
}
