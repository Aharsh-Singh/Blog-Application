package controller

import (
	"myapp/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetBlog(context *gin.Context) {
	Id := context.Param("userId")
	UserID, err := strconv.Atoi(Id)
	if err != nil || UserID < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var blogs []models.Blog
	if err := models.DB.Where("user_id = ?", UserID).Find(&blogs).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}
	if len(blogs) == 0{
		context.JSON(http.StatusNotFound, gin.H{"message": "No blogs to show"})
		return
	}
	context.JSON(http.StatusOK, blogs)
}

func GetABlog(context *gin.Context) {
	ID := context.Param("blogId")
	blogID, err := strconv.Atoi(ID)
	if err != nil || blogID < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Blog Id"})
		return
	}
	var blog models.Blog
	if err := models.DB.Where("ID = ?", blogID).Find(&blog).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	context.JSON(http.StatusOK, blog)
}

func CreateBlog(context *gin.Context) {
	Id := context.Param("userId")
	userID, err := strconv.Atoi(Id)
	if err != nil || userID < 0{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invaild userId"})
		return 
	}

	var requestBody struct{
		Title string
		Content string
	}

	if err:= context.BindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invaild input"})
		return
	}

	blog := models.Blog{
		Title : requestBody.Title,
		Content : requestBody.Content,
		UserID : uint(userID),
	}

	if err := models.DB.Create(&blog).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Data not stored"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"blog":    blog,
	})
}

func DeleteBlog(context *gin.Context) {
	ID := context.Param("blogId")
	blogID, err := strconv.Atoi(ID)
	if err != nil || blogID < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invaild blog Id"})
		return
	}
	var blog models.Blog
	if err := models.DB.Delete(&blog, blogID).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error in deleting blog"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Blog deleted successfully",
	})
}

func UpdateABlog(context *gin.Context) {
	blogId := context.Param("blogId")
	id, err := strconv.Atoi(blogId)
	if err != nil || id < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	var blog models.Blog
	if err := models.DB.First(&blog, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	var patchBlog map[string]string
	if err := context.BindJSON(&patchBlog); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if title, ok := patchBlog["title"]; ok {
		blog.Title = title
	}
	if content, ok := patchBlog["content"]; ok {
		blog.Content = content
	}

	if err := models.DB.Save(&blog).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}

	context.JSON(http.StatusOK, blog)
}
