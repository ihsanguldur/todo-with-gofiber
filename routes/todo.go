package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo/handlers"
	"todo/middlewares"
)

func TodoRouter(api fiber.Router) {
	api.Get("/todo/:user_id", middlewares.Protected(), handlers.GetUserTodos)
	api.Post("/todo/:user_id", middlewares.Protected(), handlers.CreateTodo)
	api.Put("/todo/:user_id", middlewares.Protected(), handlers.UpdateTodo)
	api.Delete("/todo", middlewares.Protected(), handlers.DeleteTodo)
}
