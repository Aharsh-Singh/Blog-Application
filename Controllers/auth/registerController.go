package auth

import (
	"myapp/models"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	supabase_auth "github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
)

func UserRegister(context *gin.Context){
	var requestBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectReference := os.Getenv("PROJECT_REFERENCE")
	apiKey := os.Getenv("API_KEY")
	print(projectReference)
	client := supabase_auth.New(projectReference, apiKey)

	resp, err := client.Signup(types.SignupRequest{
		Email:    requestBody.Email,
		Password: requestBody.Password,
	})
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Supabase_ID:    resp.User.ID.String(),
		Email: resp.User.Email,
		Name:  requestBody.Name,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":       "Sign up successful!",
		"user":          user,
	})
}