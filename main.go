package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"todo/database"
	"todo/routes"
	"todo/utils"
)

func main() {

	database.ConnectDB()

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))

}
