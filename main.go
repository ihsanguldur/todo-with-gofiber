package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"todo/database"
	"todo/routes"
)

func main() {

	app := fiber.New()
	database.ConnectDB()

	api := app.Group("/api")
	routes.UserRouter(api)
	routes.TodoRouter(api)

	log.Fatal(app.Listen(":8080"))

}
