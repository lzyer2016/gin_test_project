package main

import "github.com/gin-gonic/gin"

func main()  {

	r := gin.Default()
	r.GET("/get_param", func(context *gin.Context) {
		id := context.Query("id")
		name := context.DefaultQuery("name", "hello")
		context.String(200, id+"::"+name)
	})
	r.Run(":8080")

}
