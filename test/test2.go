package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "upload books!",
		})
	})
	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete books!",
		})
	})
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "upload books!",
		})
	})
	r.Run(":9091")
}
