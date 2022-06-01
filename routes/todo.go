package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo/handlers"
	"todo/middlewares"
)

func TodoRouter(api fiber.Router) {
	api.Get("/todo", middlewares.Protected(), handlers.GetUserTodos)
	api.Post("/todo", middlewares.Protected(), handlers.CreateTodo)
	api.Put("/todo", middlewares.Protected(), handlers.UpdateTodo)
	api.Delete("/todo", middlewares.Protected(), handlers.DeleteTodo)
}
