package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 参数绑定
func main() {
	type user struct {
		Username string `form:"username" json:"uname"`
		Password string `form:"password" json:"pwd"`
	}
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		var u user
		err := c.ShouldBind(&u) // 按值传递
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":   "请求的参数不正确!",
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok!",
				"user":    u,
			})
		}

	})
	r.POST("/user", func(c *gin.Context) {
		var u user
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":   "请求的参数不正确!",
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok!",
				"user":    u,
			})
		}

	})

	r.Run(":9001")
}
