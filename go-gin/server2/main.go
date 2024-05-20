package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var cli *redis.Client

// Gin web
// bind :8081
// one GET method
func main() {
	cli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := cli.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		st := time.Now()

		deadline, ok := c.Deadline()
		log.Printf("deadline: %v, ok: %v\n", deadline, ok)
		deadline, ok = c.Request.Context().Deadline()
		log.Printf("deadline: %v, ok: %v\n", deadline, ok)

		ch := time.After(1 * time.Second)
		select {
		case <-ch:
		case <-c.Request.Context().Done():
		}

		res, err := callRedis(c)
		if err != nil {
			// log error and time cost
			log.Printf("time cost: %v, err: %v\n", time.Since(st), err)
			return
		}

		log.Printf("time cost: %v, res: %v\n", time.Since(st), res)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
}

func callRedis(ctx context.Context) (string, error) {
	return cli.Get(ctx, "name").Result()
}
