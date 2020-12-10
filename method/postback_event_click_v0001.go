package method

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"wx_api/common"
)

func EventClickV0001(ctx *gin.Context, data []byte) {
	fmt.Printf("data: %+v \n", string(data))
	eventSubscribe, err := common.ParseEventClick(data)
	if err != nil {
		fmt.Sprintf("parse event click v0001 err:%+v", string(data))
		return
	}
	// 构造返回xml
	respMsg := common.RespMsgText{
		Content: "i am panda",
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
