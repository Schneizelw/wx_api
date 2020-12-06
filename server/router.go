package server

import (
	"github.com/gin-gonic/gin"
	"wx_api/method"
)

func router(r *gin.Engine) {
	// 测试接口
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// common
	common := r.Group("common")
	{
		common.POST("/token", method.AccessToken)
	}

}
