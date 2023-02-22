package models

import (
	"time"

	_ "gorm.io/gorm"

	_ "github.com/go-playground/validator/v10"
)

type ToDo struct {
	ID          uint `gorm:"primarykey"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserID      int    `json:"user_id" validate:"required"`
	CreatedAt time.Time
  	UpdatedAt time.Time
}
