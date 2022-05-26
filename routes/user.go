package routes

import "github.com/gofiber/fiber/v2"

func UserRouter(api fiber.Router) {

	api.Get("/user", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("user get."))
	})

	api.Post("/user", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("user post."))
	})

	api.Put("/user", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("user put."))
	})

	api.Delete("/user", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("user delete."))
	})
}
