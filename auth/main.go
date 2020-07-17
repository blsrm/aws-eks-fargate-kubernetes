package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
	"github.com/hom-bahrani/golang-gke-firestore/auth/routes"
)

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())
	app.Use(helmet.New())

	routes.SetupRoutes(app)
	fmt.Println("Listneing on Port 3000")
	_ = app.Listen(3000)
}
