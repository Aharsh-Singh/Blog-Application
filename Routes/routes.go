package routes

import (
	"myapp/controllers"
	"github.com/gin-gonic/gin"
	authController "myapp/controllers/auth"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("blogs", controllers.GetBlog)
		api.POST("blogs", controllers.CreateBlog)
		api.GET("blogs/:blogId", controllers.GetABlog)
		api.PATCH("blogs/:blogId", controllers.UpdateABlog)
		api.DELETE("blogs/:blogId", controllers.DeleteBlog)
	}
	auth := r.Group("/auth")
	{
		auth.POST("signup", authController.UserRegister)
		auth.POST("login", authController.Login)
	}
	return r
}