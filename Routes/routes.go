package Routes

import (
	"myapp/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("blogs", Controllers.GetBlog)
		api.POST("blogs", Controllers.CreateBlog)
		api.GET("blogs/:blogId", Controllers.GetABlog)
		api.PATCH("blogs/:blogId", Controllers.UpdateABlog)
		api.DELETE("blogs/:blogId", Controllers.DeleteBlog)
	}
	return r
}