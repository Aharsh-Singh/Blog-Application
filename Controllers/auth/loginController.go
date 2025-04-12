package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	supabase_auth "github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
)

func Login(context *gin.Context){
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	projectReference  := os.Getenv("PROJECT_REFERENCE")
	apiKey := os.Getenv("API_KEY")

	client := supabase_auth.New(
        projectReference,
        apiKey,
    )

	resp, err := client.Token(types.TokenRequest{
        GrantType: "password",
        Email: requestBody.Email,
        Password: requestBody.Password,
    })


    if err != nil {
        context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
    }

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
	})
}