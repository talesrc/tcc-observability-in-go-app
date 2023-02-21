package main

import (
	"fmt"
	"tcc-observability-in-go-app/configs"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(202, gin.H{
			"message": "pong",
		})
	})

	apiPort := configs.GetConfig().ApiConfig.Port
	server.Run(fmt.Sprintf(":%s", apiPort))
}
