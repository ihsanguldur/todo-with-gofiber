package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	errCode := fiber.StatusInternalServerError
	errMessage := "Something gone wrong."

	if e, ok := err.(*fiber.Error); ok {
		errCode = e.Code
		errMessage = e.Message
	}

	return ctx.Status(errCode).JSON(fiber.Map{
		"success": false,
		"message": errMessage,
		"data":    nil,
	})
}
