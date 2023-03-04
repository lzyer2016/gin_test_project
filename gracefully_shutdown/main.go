package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/shutdown", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "success")
	})

	serv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s \n", err.Error())
		}
	}()

	// 用于接受 kill -15的信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞等待信号
	<-quit
	log.Println("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown err :", err)
	}
	log.Println("server exit")
}
