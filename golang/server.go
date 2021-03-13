package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 起動
	server()
}

func server() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world from api",
		})
	})
	r.Run(":8080")
}
