package main

import (
	"tcc-observability-in-go-app/toDoService/pkg/models"
	"tcc-observability-in-go-app/toDoService/pkg/repositories"
)

func main() {
	repositories.ConnectToDB()
	repositories.DB.AutoMigrate(&models.ToDo{})
}
