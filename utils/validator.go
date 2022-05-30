package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"net/mail"
	"strconv"
)

func IsPasswordValid(password string) bool {
	return len(password) >= 6
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsJWTValid(t *jwt.Token, i string) bool {
	id, err := strconv.Atoi(i)

	if err != nil {
		return false
	}

	uid := int(t.Claims.(jwt.MapClaims)["user_id"].(float64))

	if id != uid {
		return false
	}

	return true
}
