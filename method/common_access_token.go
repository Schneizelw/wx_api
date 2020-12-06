package method

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guonaihong/gout"
	"net/http"
	"wx_api/common"
)

func AccessToken(ctx *gin.Context) {
	url := fmt.Sprintf(common.WxApiCommonAccessToken, common.WxAppId, common.WxAppSecret)
	resp := common.WxRespAccessToken{}
	err := gout.GET(url).BindJSON(&resp).Do()
	if err != nil {
		fmt.Printf("access token err: %s \n", err)
		return
	}
	fmt.Printf("resp: %+v", resp)
	ctx.JSON(http.StatusOK, gin.H{
		"errNum": common.ErrNumSuccess,
		"errMsg": common.ErrMsg(common.ErrNumSuccess),
	})
	return
}
