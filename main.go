package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello, World!"})
    })
	//fmt.Println("hello word")
    r.Run()
}
