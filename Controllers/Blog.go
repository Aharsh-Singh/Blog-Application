package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

var blogs = []map[string]string{
	{"id": "1", "title": "Go Basics", "content": "An introduction to Go programming language."},
	{"id": "2", "title": "Building APIs with Gin", "content": "How to create RESTful APIs using Gin framework."},
	{"id": "3", "title": "Concurrency in Go", "content": "Understanding goroutines and channels in Go."},
	{"id": "4", "title": "Go vs Python", "content": "Comparing Go and Python for backend development."},
	{"id": "5", "title": "Deploying Go Apps", "content": "A step-by-step guide to deploying Go applications."},
}


func GetBlog(context *gin.Context) {
	if len(blogs) == 0{
		context.JSON(http.StatusNotFound, gin.H{"message": "No blogs to show"})
		return
	}
	context.JSON(http.StatusOK, blogs)
}

func GetABlog(context *gin.Context) {
	blogID := context.Param("blogId")
	for index := range blogs {
		if blogs[index]["id"] == blogID {
			context.JSON(http.StatusOK, blogs[index])
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
}

func CreateBlog(context *gin.Context) {
	var newBlog map[string]string
	if err := context.BindJSON(&newBlog); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newBlog["id"] = strconv.Itoa(len(blogs)+1)
	blogs = append(blogs, newBlog)
	context.JSON(http.StatusCreated, newBlog)
}

func DeleteBlog(context *gin.Context) {
	blogID := context.Param("blogId")
	for index, blog := range blogs {
		if blog["id"] == blogID {
			blogs = append(blogs[:index], blogs[index+1:]...)
			context.JSON(http.StatusOK, gin.H{"message": "Blog deleted"})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
}

func UpdateABlog(context *gin.Context){
	var patchBlog map[string]string
	blogId := context.Param("blogId")
	if err := context.BindJSON(&patchBlog); err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	for index := range blogs{
		if blogs[index]["id"] == blogId{
			value, exists := patchBlog["title"]
			if exists{
				blogs[index]["title"] = value 
			}

			value, exists = patchBlog["content"]
			if exists{
				blogs[index]["content"] = value
			}
			context.JSON(http.StatusOK, blogs[index])
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
}