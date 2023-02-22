package main

import (
	"fmt"
	"tcc-observability-in-go-app/toDoService/configs"
	"tcc-observability-in-go-app/toDoService/pkg/repositories"
	"tcc-observability-in-go-app/toDoService/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	repositories.ConnectToDB()
}

func main() {
	server := gin.Default()

	server.POST("/todo", controllers.CreateToDo)

	apiPort := configs.GetConfig().ApiConfig.Port
	server.Run(fmt.Sprintf(":%s", apiPort))
}
