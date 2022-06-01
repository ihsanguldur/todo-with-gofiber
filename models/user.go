package models

import (
	"time"
)

type User struct {
	UserID       uint      `gorm:"primaryKey" json:"user_id"`
	UserName     string    `gorm:"not null" json:"user_name"`
	UserSurname  string    `json:"user_surname"`
	UserEmail    string    `json:"user_email"`
	UserPassword string    `json:"user_password"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
