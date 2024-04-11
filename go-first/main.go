package main

// Use Gin to build a http server
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
