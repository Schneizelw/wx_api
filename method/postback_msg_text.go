package method

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"wx_api/common"
)

func MsgText(ctx *gin.Context, data []byte) {
	// 解析请求xml
	fmt.Printf("data: %+v \n", string(data))
	userMsg, err := common.ParseMsgText(data)
	if err != nil {
		fmt.Sprintf("parse msg type err:%+v", string(data))
		return
	}
	// 构造返回xml
	respMsg := common.RespMsgText{
		Content: userMsg.Content,
	}
	respMsg.ToUserName = userMsg.FromUserName
	respMsg.FromUserName = userMsg.ToUserName
	respMsg.CreateTime = string(time.Now().Unix())
	respMsg.MsgType = userMsg.MsgType
	resp, err := common.ParseToXml(respMsg)
	if err != nil {
		fmt.Sprintf("parse to xml err: %+v", err)
		return
	}
	fmt.Printf("resp: %+v", resp)
	ctx.Writer.WriteString(resp)
	return
}
