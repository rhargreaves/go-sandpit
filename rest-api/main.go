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

	r.GET("/json", func(c *gin.Context) {
		// gin.H is an alias for map[string]interface{}
		c.JSON(200, gin.H{"message": "Hello, Docker with Go!"})
	})

	r.Run(":8080")
}
