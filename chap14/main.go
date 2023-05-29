package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 这也是一个中间件
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"index": "Ok",
	})
}

// 定义一个新的中间件
func m1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"m1": "before!",
	})

	start := time.Now()
	c.Next()
	// 停止调用后续的函数
	// c.Abort()
	cost := time.Since(start)
	fmt.Printf("花费时间:%v\n", cost)
	c.JSON(http.StatusOK, gin.H{
		"m1": "after!",
	})

}
func main() {
	r := gin.Default()
	r.Use(m1) // 全局注册中间件
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"shop": "welcome!",
		})
	})
	r.Run(":9001")

}
