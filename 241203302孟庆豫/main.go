package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./package.html", "./target.html")
	r.GET("/package.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "package.html", "nil")
	})
	r.GET("/target.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "target.html", "nil")
	})

	_ = r.Run("127.0.0.1:8080")
}
