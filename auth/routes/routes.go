package routes

import (
	"github.com/gofiber/fiber"

	"github.com/hom-bahrani/golang-gke-firestore/auth/controllers"
)

// SetupRoutes creates the initial routes for the auth service
func SetupRoutes(app *fiber.App) {
	app.Get("/api/users/ping", controllers.Ping)
	app.Get("/api/users/currentuser/:id", controllers.CurrentUser)
	app.Post("/api/users/signup", controllers.Signup)
	app.Post("/api/users/signin", controllers.Signin)
	app.Post("/api/users/signout", controllers.Signout)
}
