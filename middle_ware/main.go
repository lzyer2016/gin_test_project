package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/log", func(c *gin.Context) {
		// panic("test panic")
		c.String(http.StatusOK, "hello world")
	})
	r.Run(":8080")
}
