package controller

import (
	"myapp/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	context.JSON(http.StatusOK, users)
}