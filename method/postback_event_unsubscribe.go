package method

import (
	"github.com/gin-gonic/gin"
)

func EventUnsubscribe(ctx *gin.Context, data []byte) {
	/*
		fmt.Printf("data: %+v \n", string(data))
		eventUnsubscribe, err := common.ParseEventUnsubscribe(data)
		if err != nil {
			fmt.Sprintf("parse event subscribe err:%+v", string(data))
			return
		}
		// 构造返回xml
		respMsg := common.RespMsgText{
			Content: "不要想不开，麻烦重新关注一下谢谢",
		}
		respMsg.ToUserName = eventUnsubscribe.FromUserName
		respMsg.FromUserName = eventUnsubscribe.ToUserName
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
	*/
}
