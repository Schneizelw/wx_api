package method

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"wx_api/common"
)

func EventSubscribe(ctx *gin.Context, data []byte) {
	fmt.Printf("data: %+v \n", string(data))
	eventSubscribe, err := common.ParseEventSubscribe(data)
	if err != nil {
		fmt.Sprintf("parse event subscribe err:%+v", string(data))
		return
	}
	// 构造返回xml
	respMsg := common.RespMsgText{
		Content: "你发现了你不该发现的公众号",
	}
	respMsg.ToUserName = eventSubscribe.FromUserName
	respMsg.FromUserName = eventSubscribe.ToUserName
	respMsg.CreateTime = string(time.Now().Unix())
	respMsg.MsgType = common.MsgText
	resp, err := common.ParseToXml(respMsg)
	if err != nil {
		fmt.Sprintf("parse to xml err: %+v", err)
		return
	}
	fmt.Printf("resp: %+v", resp)
	ctx.Writer.WriteString(resp)
	return
}
