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
		users := api.Group("/users")
		{
			users.GET("/get_users", controller.GetAllUsers)
		}

		blogs := api.Group("/blogs")
		{
			blogs.GET("get_blogs/:userId", controller.GetBlog)
			blogs.POST("write/:userId", controller.CreateBlog)
			blogs.GET("get_blog/:blogId", controller.GetABlog)
			blogs.PATCH("update/:blogId", controller.UpdateABlog)
			blogs.DELETE("delete/:blogId", controller.DeleteBlog)
		}

		auth := api.Group("/auth")
		{
			auth.POST("signup", authController.UserRegister)
			auth.POST("login", authController.Login)
		}
	}
	return r
}