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
		"user_id":       credential.UserID,
		"user_name":     credential.UserName,
		"user_surname":  credential.UserSurname,
		"user_email":    credential.UserEmail,
		"user_password": credential.UserPassword,
		"createdAt":     credential.CreatedAt,
		"updatedAt":     credential.UpdatedAt,
		"exp":           time.Now().Add(30 * time.Minute).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token, err := t.SignedString([]byte(config.Config("JWT_SECRET")))

	if err != nil {
		log.Fatal(err.Error())
	}
	
	return token
}

/*func ParseToken(t *jwt.Token) jwt.MapClaims {
	return t.Claims.(jwt.MapClaims)
}*/
