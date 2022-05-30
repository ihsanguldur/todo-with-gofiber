package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo/handlers"
	"todo/middlewares"
	"todo/utils"
)

func TodoRouter(api fiber.Router) {

	api.Get("/todo/:user_id", middlewares.Protected(), handlers.GetUserTodos)

	api.Post("/todo", func(ctx *fiber.Ctx) error {
		return utils.SuccessPresenter(nil, "todo post.", ctx)
	})

	api.Put("/todo", func(ctx *fiber.Ctx) error {
		return utils.SuccessPresenter(nil, "todo put.", ctx)
	})

	api.Delete("/todo", func(ctx *fiber.Ctx) error {
		return utils.SuccessPresenter(nil, "todo delete.", ctx)
	})
}
