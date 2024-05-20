package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Gin implements a HTTP web server
// router has one GET handler
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		// call localhost:8081 to get data.
		// use http request with GET method
		// set timeout 500ms
		st := time.Now()

		subCtx, cancel := context.WithTimeout(c, 500*time.Millisecond)
		defer cancel()

		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8081/ping", nil)
		req = req.WithContext(subCtx)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			// log error and time cost
			log.Printf("error: %v, time cost: %v\n", err, time.Since(st))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "timeout",
			})
			return
		}
		defer resp.Body.Close()

		// log response body and time cost
		log.Printf("response: %v, time cost: %v\n", resp.Status, time.Since(st))

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
