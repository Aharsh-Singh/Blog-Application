package auth

import(
	"net/http"
	"myapp/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"myapp/utils"
)

func Login(context *gin.Context){
	var requestBody struct {
		Email    string
		Password string
	}
	user := models.User{}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := models.DB.Where("email = ?", requestBody.Email).First(&user).Error; err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	Payload := utils.JWTClaims{
		ID:    user.ID,
	}
	token, err := utils.GenerateJWT(Payload)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"user":    user,
		"token":   token,
	})
}