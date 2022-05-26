package routes

import "github.com/gofiber/fiber/v2"

func TodoRouter(api fiber.Router) {

	api.Get("/todo", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("todo get."))
	})

	api.Post("/todo", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("todo post."))
	})

	api.Put("/todo", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("todo put."))
	})

	api.Delete("/todo", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("todo delete."))
	})
}
