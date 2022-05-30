package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm/clause"
	"strconv"
	"todo/database"
	"todo/models"
	"todo/utils"
)

func CreateTodo(ctx *fiber.Ctx) error {
	var err error
	id := ctx.Params("user_id")
	todo := new(models.Todo)
	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsJWTValid(token, id) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token id.")
	}

	if err = ctx.BodyParser(todo); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted body.")
	}

	uid, _ := strconv.ParseUint(id, 10, 64)
	todo.UserID = uint(uid)

	if err = database.DB.Create(todo).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Database error.")
	}

	return utils.SuccessPresenter(todo, fmt.Sprintf("%s", id), ctx)
}

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

func UpdateTodo(ctx *fiber.Ctx) error {
	var err error
	todo := &models.Todo{}
	id := ctx.Params("user_id")
	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsJWTValid(token, id) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token id.")
	}

	if err = ctx.BodyParser(todo); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted body.")
	}

	result := database.DB.Where("todo_id = ?", todo.TodoID).Updates(todo)

	if err = result.Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "database error.")
	}

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "todo not found.")
	}

	return utils.SuccessPresenter(todo, "Todo updated.", ctx)
}

func DeleteTodo(ctx *fiber.Ctx) error {
	var err error
	todo := new(models.Todo)
	uid := ctx.Query("uid")
	tid := ctx.Query("tid")
	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsJWTValid(token, uid) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token id.")
	}

	if uid == "" || tid == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Wrong query notation.")
	}

	result := database.DB.Clauses(clause.Returning{}).Delete(todo, tid)

	if err = result.Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "database err")
	}

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "todo not found.")
	}

	return utils.SuccessPresenter(todo, "Todo deleted.", ctx)
}
