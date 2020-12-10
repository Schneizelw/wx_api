package method

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wx_api/common"
)

func CheckServer(ctx *gin.Context) {
	opt := &common.CheckSignatureOpt{
		Signature: ctx.Query("signature"),
		Timestamp: ctx.Query("timestamp"),
		Nonce:     ctx.Query("nonce"),
	}
	if !common.CheckSignature(opt) {
		fmt.Println("error signature")
		return
	}
	ctx.Writer.WriteString(ctx.Query("echostr"))
	return
}
