package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"todo/database"
)

func main() {

	app := fiber.New()
	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		
		return c.Send([]byte("Hello"))

	})

	log.Fatal(app.Listen(":8080"))

}
