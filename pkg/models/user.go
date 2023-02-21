package models

import (
	"gorm.io/gorm"
	"regexp"
)

type User struct {
	gorm.Model
	ID        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *User) IsCPFValid() (matched bool) {
	matched, _ = regexp.MatchString(`^\d{3}\.\d{3}\.\d{3}\-\d{2}$`, u.CPF)
	return
}
