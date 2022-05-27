package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo/handlers"
)

func UserRouter(api fiber.Router) {

	api.Post("/user", handlers.Create)
	api.Get("/user/:id", handlers.GetUser)
}
