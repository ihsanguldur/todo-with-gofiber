package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
	"todo/config"
	"todo/models"
)

func GenerateToken(credential *models.User) string {
	claim := jwt.MapClaims{
		"user": credential,
		"exp":  time.Now().Add(15 * time.Second).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token, err := t.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		log.Fatal(err.Error())
	}
	return token
}
