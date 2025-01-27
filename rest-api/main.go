package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Header("Another-Header", "blah")
		c.String(200, "Hello, Docker with Go!\n")
	})
	r.Run(":8080")
}
