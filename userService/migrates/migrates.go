package main

import (
	"tcc-observability-in-go-app/pkg/models"
	"tcc-observability-in-go-app/pkg/repositories"
)

func main() {
	repositories.ConnectToDB()
	repositories.DB.AutoMigrate(&models.User{})
}
