package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/get", func(context *gin.Context) {
		context.String(200, "get")
	})
	r.GET("/post", func(context *gin.Context) {
		context.String(200, "post")
	})
	r.GET("/user/*action", func(context *gin.Context) {
		context.String(200, "hello user")
	})
	r.Handle("DELETE", "/delete", func(context *gin.Context) {
		context.String(200, "delete")
	})
	r.Run()
}
