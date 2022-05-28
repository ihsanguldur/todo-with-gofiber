package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo/handlers"
	"todo/middlewares"
)

func UserRouter(api fiber.Router) {

	api.Post("/user", handlers.Create)
	api.Get("/user/:id", handlers.GetUser)
	api.Get("/user", handlers.GetUsers)
	api.Put("/user", middlewares.Protected(), handlers.UpdateUser)
	api.Delete("/user/:id", handlers.DeleteUser)
}
