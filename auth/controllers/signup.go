package controllers

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber"
	models "github.com/hom-bahrani/golang-gke-firestore/auth/models"
)

const collectionName string = "users"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Signup creates a new user
func Signup(c *fiber.Ctx) {
	user := new(models.User)

	rand.Seed(time.Now().UnixNano())
	user.ID = randSeq(10)

	_ = c.JSON(&fiber.Map{
		"success": true,
		"data":    user,
	})
}
