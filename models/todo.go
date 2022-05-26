package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	ID        uint      `json:"todo_id"`
	Body      string    `json:"todo_body"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
