package main

// Use Gin to build a http server
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func main() {
	// ginTest()

	chanTest()
}

func ginTest() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}

func chanTest() {
	ch := make(chan int)
	ch <- 1
	fmt.Println("1")
	ch <- 2
	fmt.Println("2")
}
