package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/post_body", func(context *gin.Context) {
		bodyBuf, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			context.Abort()
		}
		context.String(http.StatusOK, string(bodyBuf))
	})

	r.GET("/test_bind", testing)
	r.POST("/test_bind", testing)

	r.Run(":8080")
}

type Person struct {
	Id   int    `form:"id" binding:"required,gt=10"`
	Name string `form:"name" binding:"required"`
	Addr string `form:"addr"`
}

func testing(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err == nil {
		c.String(http.StatusOK, "%v", person)
	} else {
		c.String(http.StatusBadRequest, "person bind error %v", err)
	}
}
