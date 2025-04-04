package auth

import (
	"net/http"
	"myapp/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(context *gin.Context){
	user := models.User{}

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser := models.User{}
	if err := models.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error in hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := models.DB.Create(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}