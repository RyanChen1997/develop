package main

// start a gin server
// register one GET function
// return 'hello world'

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func server() {
	// signal
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// register Gin
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		go func(done <-chan struct{}) {
			<-done
			log.Println("return bye")
		}(mainCtx.Done())

		time.Sleep(10 * time.Second)
		log.Println("return hello")
		c.String(http.StatusOK, "hello world")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	<-mainCtx.Done()
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
