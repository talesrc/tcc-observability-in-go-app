package controllers

import (
	"net/http"
	"tcc-observability-in-go-app/toDoService/pkg/models"
	"tcc-observability-in-go-app/toDoService/pkg/repositories"

	"github.com/gin-gonic/gin"
)

func CreateToDo(c *gin.Context) {
	var newToDo models.ToDo
	c.BindJSON(&newToDo)

	repositories.DB.Create(&newToDo)
	c.JSON(http.StatusCreated, gin.H{
		"status_code": http.StatusCreated,
		"message": "The to do was created successfully",
	})
}

