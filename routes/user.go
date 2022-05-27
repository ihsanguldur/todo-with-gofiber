package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo/utils"
)

func UserRouter(api fiber.Router) {

	api.Get("/user", func(ctx *fiber.Ctx) error {
		return utils.SuccessPresenter(nil, "user get.", ctx)
	})

	api.Post("/user", func(ctx *fiber.Ctx) error {
		return utils.SuccessPresenter(nil, "user post.", ctx)
	})

	api.Put("/user", func(ctx *fiber.Ctx) error {
		return utils.SuccessPresenter(nil, "user put.", ctx)
	})

	api.Delete("/user", func(ctx *fiber.Ctx) error {
		return utils.SuccessPresenter(nil, "user delete.", ctx)
	})
}
