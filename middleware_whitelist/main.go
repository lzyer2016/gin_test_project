package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IpAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{"127.0.0.2"}
		flag := false
		clientIp := c.ClientIP()
		for _, ip := range ipList {
			if clientIp == ip {
				flag = true
			}
		}
		if !flag {
			c.String(401, "%s not in iplist", clientIp)
			c.Abort()
		}
	}
}
func main() {
	r := gin.Default()
	r.Use(IpAuthMiddleware())
	r.GET("/request", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
	r.Run(":8080")
}
