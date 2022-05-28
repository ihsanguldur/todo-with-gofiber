package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"todo/database"
	"todo/models"
	"todo/utils"
)

func Login(ctx *fiber.Ctx) error {

	type LoginInput struct {
		Email    string `json:"user_email"`
		Password string `json:"user_password"`
	}

	var err error
	var inputs LoginInput
	user := new(models.User)

	if err = ctx.BodyParser(&inputs); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted body.")
	}

	if !utils.IsPasswordValid(inputs.Password) {
		return fiber.NewError(fiber.StatusBadRequest, "Password must be 6 characters.")
	}

	if !utils.IsEmailValid(inputs.Email) {
		return fiber.NewError(fiber.StatusBadRequest, "Email is not valid.")
	}

	if err = database.DB.Where("user_email = ?", inputs.Email).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Email is not correct.")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong while finding.")
	}

	if utils.ComparePassword(inputs.Password, user.UserPassword) != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Password is not correct.")
	}

	token := utils.GenerateToken(user)

	return utils.SuccessPresenter(fiber.Map{
		"user":  user,
		"token": token,
	}, "Login successful.", ctx)
}
