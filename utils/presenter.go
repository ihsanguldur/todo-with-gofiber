package utils

import "github.com/gofiber/fiber/v2"

func SuccessPresenter(data interface{}, message string, ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})

}

func ErrorPresenter(errCode int, errMessage string, ctx *fiber.Ctx) error {

	return ctx.Status(errCode).JSON(fiber.Map{
		"success": false,
		"message": errMessage,
		"data":    nil,
	})

}
