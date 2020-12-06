package main

import (
	"github.com/gin-gonic/gin"
	"wx_api/server/router.go"
)

func main() {
	r := gin.Default()
	router.
	r.Run("127.0.0.1:5021")
}
