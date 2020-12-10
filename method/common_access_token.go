package method

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wx_api/common"
)

func AccessToken(ctx *gin.Context) {
	token, _ := common.GetAccessToken()
	ctx.JSON(http.StatusOK, gin.H{
		"AccessToken": token,
		"errNum":      common.ErrNumSuccess,
		"errMsg":      common.ErrMsg(common.ErrNumSuccess),
	})
	return
}
