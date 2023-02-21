package main

import (
	"fmt"
	"tcc-observability-in-go-app/configs"
	"tcc-observability-in-go-app/pkg/controllers"
	"tcc-observability-in-go-app/pkg/repositories"

	"github.com/gin-gonic/gin"
)

func init() {
	repositories.ConnectToDB()
}

func main() {
	server := gin.Default()

	server.POST("/user", controllers.CreateUser)
	server.GET("/user/:id", controllers.GetUserById)

	apiPort := configs.GetConfig().ApiConfig.Port
	server.Run(fmt.Sprintf(":%s", apiPort))
}
