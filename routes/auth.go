package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo/handlers"
)

func AuthRouter(api fiber.Router) {

	api.Post("auth/login", handlers.Login)

}
