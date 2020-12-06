package server

import (
	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()
	router(r)
	r.Run("127.0.0.1:5021")
}
