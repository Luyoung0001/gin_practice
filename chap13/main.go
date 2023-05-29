package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 访问 /index 这个路由
	// 获取信息
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	// 提交
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})

	})
	// 更新数据
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})

	})
	// 删除数据
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})

	})
	// 请求方法大杂烩
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case "PUT":
			c.JSON(http.StatusOK, gin.H{
				"method": "PUT",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"method": "Any",
		})

	})
	// 针对意料之外的访问地址,没有这个路由时
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "你似乎来到了陌生的地带~",
		})

	})
	// 路由组的概念
	// 把公用的前缀提取出来,创建一个路由组
	// 路由组是可以嵌套定义的,也就说,路由组的定义是递归的
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})

		})
		videoGroup.GET("/xxx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video/xxx",
			})

		})
		videoGroup.GET("/ooo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video/ooo",
			})

		})
	}

	err := r.Run(":9001")
	if err != nil {
		return
	}
}
