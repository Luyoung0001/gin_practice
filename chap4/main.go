package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("/Users/luliang/GoLand/gin_practice/chap4/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "你好,前端真是太有意思了!",
		})
	})
	r.Run(":9091")
}
