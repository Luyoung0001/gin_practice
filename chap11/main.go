package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("/Users/luliang/GoLand/gin_practice/chap11/index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)

	})
	r.POST("/upload", func(c *gin.Context) {
		// 读取文件
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// 将读取的文件保存到本地
			// dsf := fmt.Sprintf("./%s", f.Filename)
			dsf := path.Join("./", f.Filename)
			err = c.SaveUploadedFile(f, dsf)
			if err != nil {
				return
			} // 将文件 f 保存到 dsf 中
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})

		}

	})
	r.Run(":9001")
}
