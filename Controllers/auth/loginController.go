package auth

import(
	"net/http"
	// "myapp/models"
	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
	// "myapp/utils"
	authgo "github.com/supabase-community/auth-go"
	"os"
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

	client := authgo.New(
        projectReference,
        apiKey,
    )

    resp, err := client.Token(authgo.TokenGrantRequest{
        GrantType: "password",
        Email: requestBody.Email,
        Password: requestBody.Password,
    })
    if err != nil {
        context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
    }

	// Payload := utils.JWTClaims{
	// 	ID:    user.ID,
	// }
	// token, err := utils.GenerateJWT(Payload)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	// 	return
	// }

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
	})
}