package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Send([]byte("Hello"))

	})

	log.Fatal(app.Listen(":8080"))

}
