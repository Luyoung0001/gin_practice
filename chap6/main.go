package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	type usr struct {
		Name string `json:"name"`
		Msg  string
		Age  int
	}
	r.GET("/json", func(c *gin.Context) {
		data := gin.H{
			"name":    "小王子!",
			"massage": "hello",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})
	r.GET("/another_json", func(c *gin.Context) {
		data := usr{
			"小王子!",
			"hello,golang!",
			18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9001")
}
