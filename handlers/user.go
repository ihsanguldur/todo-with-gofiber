package handlers

import (
	"github.com/gofiber/fiber/v2"
	"todo/database"
	"todo/models"
	"todo/utils"
)

func Create(ctx *fiber.Ctx) error {

	var err error
	user := new(models.User)

	if err = ctx.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong while parsing.")
	}

	if err = database.DB.Create(user).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please Check Your Credentials.")
	}

	return utils.SuccessPresenter(user, "User created.", ctx)
}
