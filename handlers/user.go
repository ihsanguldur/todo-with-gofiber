package handlers

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"todo/database"
	"todo/models"
	"todo/utils"
)

func Create(ctx *fiber.Ctx) error {

	var err error
	user := new(models.User)

	if err = ctx.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something Went Wrong While Parsing.")
	}

	if !utils.IsPasswordValid(user.UserPassword) {
		return fiber.NewError(fiber.StatusBadRequest, "Password must be 6 character.")
	}

	if !utils.IsEmailValid(user.UserEmail) {
		return fiber.NewError(fiber.StatusBadRequest, "Email is not valid.")
	}

	user.UserPassword = utils.HashPassword(user.UserPassword)

	if err = database.DB.Create(user).Error; err != nil {
		message := "Please Check Your Credentials."
		if strings.Contains(err.Error(), "23505") {
			message = "Email is already in use."
		}
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	return utils.SuccessPresenter(user, "User created.", ctx)
}
