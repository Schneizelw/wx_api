package method

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wx_api/common"
)

func AuthRedirect(ctx *gin.Context) {
	fmt.Printf("???")
	opt := &common.OauthRedirectParam{
		Code: ctx.Query("code"),
	}
	resp, err := common.GetAccessTokenOauth2(opt.Code)
	if err != nil {
		fmt.Printf("get access token oauth2 err: %s", err.Error())
		return
	}
	userInfo, err := common.GetUserInfoOauth2(resp.AccessToken, resp.Openid)
	if err != nil {
		fmt.Printf("get user info oauth2 err: %s", err.Error())
		return
	}
	fmt.Printf("userinfo: %+v\n", userInfo)
	return
}
