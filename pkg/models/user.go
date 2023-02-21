package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID int
	FirstName string
	LastName string
	Email string
	Password string
}