package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"wx_api/common"
	"wx_api/method"
)

func router(r *gin.Engine) {
	// 测试接口
	r.GET("/wx_api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 验证服务地址
	r.GET("/wx_api/postback", method.CheckServer)

	// common
	common := r.Group("/wx_api/common")
	{
		common.GET("/token", method.AccessToken)
	}

	// menu
	menu := r.Group("/wx_api/menu")
	{
		menu.GET("/create", method.MenuCreate)
	}

	// postback接口
	r.POST("/wx_api/postback", PostbackRoute)

	// 授权redirect接口
	r.GET("/wx_api/user_info", method.AuthRedirect)
}

var PostBackRouteFunc = map[string]func(ctx *gin.Context, data []byte){
	common.MsgText:          method.MsgText,
	common.MsgImage:         method.MsgImage,
	common.EventSubscribe:   method.EventSubscribe,
	common.EventUnsubscribe: method.EventUnsubscribe,
	common.EventClickV0001:  method.EventClickV0001,
	common.EventClickV0002:  method.EventClickV0002,
}

func PostbackRoute(ctx *gin.Context) {
	// 检查是否是微信的回调
	opt := &common.CheckSignatureOpt{
		Signature: ctx.Query("signature"),
		Timestamp: ctx.Query("timestamp"),
		Nonce:     ctx.Query("nonce"),
	}
	if !common.CheckSignature(opt) {
		fmt.Println("error signature")
		return
	}
	// 获取回调消息类型
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Printf("data: %+v \n", string(data))
	msgType, err := common.ParseMsgType(data)
	fmt.Printf("msg type: %+v \n", msgType)
	if err != nil {
		fmt.Sprintf("parse msg type err:%+v", string(data))
		return
	}
	// 根据消息类型做处理
	handler, ok := PostBackRouteFunc[msgType]
	if !ok {
		fmt.Sprintf("err msg type: %+v", string(data))
		return
	}
	handler(ctx, data)
	return
}
