package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo/utils"
)

func TodoRouter(api fiber.Router) {

	api.Get("/todo", func(ctx *fiber.Ctx) error {
		return utils.SuccessPresenter(nil, "todo get.", ctx)
	})

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
