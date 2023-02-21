package controllers

import (
	"net/http"
	"tcc-observability-in-go-app/pkg/models"
	"tcc-observability-in-go-app/pkg/repositories"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser models.User
	c.BindJSON(&newUser)
	if !newUser.IsCPFValid() {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "You must give a valid CPF!",
		})
		return
	}

	repositories.DB.Create(&newUser)

	c.JSON(http.StatusCreated, gin.H{
		"status_code": http.StatusCreated,
		"message":     "The user was successfuly created!",
		"user_id":     newUser.ID,
	})
	return
}

func GetUserById(c *gin.Context) {
	userId := c.Param("id")

	var user models.User
	repositories.DB.First(&user, userId)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     "User not found!",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status_code": http.StatusAccepted,
		"user":        user,
	})
	return
}
