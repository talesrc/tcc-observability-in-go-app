package main

import (
	"tcc-observability-in-go-app/configs"
	"tcc-observability-in-go-app/pkg/models"
)

func main() {
	configs.ConnectToDB()
	configs.DB.AutoMigrate(&models.User{})
}
