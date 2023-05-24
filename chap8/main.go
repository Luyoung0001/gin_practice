package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("/Users/luliang/GoLand/gin_practice/chap8/login.html", "/Users/luliang/GoLand/gin_practice/chap8/index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		// 获取 form 表单的数据
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		// 第二种方式
		username := c.DefaultPostForm("username", "用户名")
		password := c.DefaultPostForm("password", "密码")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9001")

}
