package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint      `json:"user_id"`
	Name      string    `json:"user_name"`
	Surname   string    `json:"user_surname"`
	Email     string    `json:"user_email"`
	Password  string    `json:"user_password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
