package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"os"
	"todo/utils"
)

func Protected() fiber.Handler {

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: jwtError,
	})

}

func jwtError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return utils.ErrorPresenter(fiber.StatusBadRequest, "Missing or malformed JWT", ctx)
	}
	return utils.ErrorPresenter(fiber.StatusUnauthorized, "Invalid or expired JWT", ctx)
}
