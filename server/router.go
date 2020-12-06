package main

import "github.com/gin-gonic/gin"

func route("") {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("127.0.0.1:5021")
}
