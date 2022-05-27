package models

import (
	"time"
)

type Todo struct {
	TodoID    uint      `gorm:"primaryKey" json:"todo_id"`
	TodoBody  *string   `json:"todo_body"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    int       `json:"user_id"`
}
