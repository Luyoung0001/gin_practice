package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器发送的 query 字段
		// name := c.Query("query")
		// http://127.0.0.1:9001/web?query=杨超越&age=18
		name := c.DefaultQuery("query", "someone")
		age := c.DefaultQuery("age", "age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9001")

}
