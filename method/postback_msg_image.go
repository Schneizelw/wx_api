package method

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"wx_api/common"
)

func MsgImage(ctx *gin.Context, data []byte) {
	fmt.Printf("data: %+v \n", string(data))
	userMsg, err := common.ParseMsgImage(data)
	if err != nil {
		fmt.Sprintf("parse msg type err:%+v", string(data))
		return
	}
	respMsg := common.RespMsgImage{
		MediaId: userMsg.MediaId,
	}
	respMsg.FromUserName = userMsg.ToUserName
	respMsg.ToUserName = userMsg.FromUserName
	respMsg.MsgType = userMsg.MsgType
	respMsg.CreateTime = string(time.Now().Unix())
	resp, err := common.ParseToXml(respMsg)
	if err != nil {
		fmt.Sprintf("parse to xml err: %+v", err)
		return
	}
	fmt.Printf("resp: %+v", resp)
	ctx.Writer.WriteString(resp)
	return
}
