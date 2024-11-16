package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "nil")
	})
	_ = r.Run("127.0.0.1:8080")