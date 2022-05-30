package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"todo/database"
	"todo/models"
	"todo/utils"
)

func Create(ctx *fiber.Ctx) error {

	var err error
	user := new(models.User)

	if err = ctx.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted body.")
	}

	if !utils.IsPasswordValid(user.UserPassword) {
		return fiber.NewError(fiber.StatusBadRequest, "Password must be 6 characters.")
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

func GetUser(ctx *fiber.Ctx) error {

	var err error
	user := new(models.User)

	id := ctx.Params("id")
	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsJWTValid(token, id) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token id.")
	}

	if err = database.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "User not found.")
		}
		return err
	}

	return utils.SuccessPresenter(user, "User Found.", ctx)
}

func GetUsers(ctx *fiber.Ctx) error {
	var err error
	users := &[]models.User{}

	if err = database.DB.Find(users).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessPresenter(users, "Users Found.", ctx)
}

func UpdateUser(ctx *fiber.Ctx) error {

	var err error
	user := &models.User{}

	if err = ctx.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted body.")
	}

	if user.UserPassword != "" {
		if !utils.IsPasswordValid(user.UserPassword) {
			return fiber.NewError(fiber.StatusBadRequest, "Password must be 6 character.")
		}
		user.UserPassword = utils.HashPassword(user.UserPassword)

	}

	if user.UserEmail != "" {
		if !utils.IsEmailValid(user.UserEmail) {
			return fiber.NewError(fiber.StatusBadRequest, "Email is not valid.")
		}
	}

	result := database.DB.Model(user).Clauses(clause.Returning{}).Updates(user)

	if err = result.Error; err != nil {
		message := "Error while updating."
		if strings.Contains(err.Error(), "23505") {
			message = "Email is already in use."
		}
		return fiber.NewError(fiber.StatusInternalServerError, message)
	}

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "User not found.")
	}

	return utils.SuccessPresenter(user, "User updated successfully.", ctx)

}

func DeleteUser(ctx *fiber.Ctx) error {

	var err error
	id := ctx.Params("id")
	user := &models.User{}

	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsJWTValid(token, id) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token id.")
	}

	result := database.DB.Clauses(clause.Returning{}).Delete(user, id)

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "User not found.")
	}

	if err = result.Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error while deleting.")
	}

	return utils.SuccessPresenter(user, "User deleted successfully.", ctx)
}
