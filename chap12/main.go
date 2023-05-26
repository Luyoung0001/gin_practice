package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 重定向应该怎么做呢?
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK,gin.H{
		//	"status":"Ok",
		//})
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")

	})
	r.GET("/a", func(c *gin.Context) {
		// 跳转到 b
		c.Request.URL.Path = "/b" // 修改请求 URL
		r.HandleContext(c)        // 继续处理 c

	})
	// 处理所有的方法
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

	})

	r.Run(":9001")

}
