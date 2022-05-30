package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"todo/database"
	"todo/models"
	"todo/utils"
)

func GetUserTodos(ctx *fiber.Ctx) error {
	var err error
	todos := &[]models.Todo{}
	id := ctx.Params("user_id")

	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsJWTValid(token, id) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token id.")
	}

	result := database.DB.Where("user_id = ?", id).Find(todos)

	if err = result.Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "database error.")
	}

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "User has not todo.")
	}

	return utils.SuccessPresenter(todos, "Todos found.", ctx)

}
