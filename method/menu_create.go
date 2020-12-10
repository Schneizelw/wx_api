package method

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wx_api/common"
)

var MenuStr = `{
     "button":[
     {
          "name":"panda",
          "type":"click",
          "key":"click_v0001"
      },
      {
           "name":"bear",
           "sub_button":[
           {
               "type":"view",
               "name":"search",
               "url":"http://www.baidu.com/"
            },
            {
               "type":"click",
               "name":"赞一下我们",
               "key":"click_v0002"
            },
  			{
               "type":"view",
               "name":"授权",
               "url":"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxa03e08dedbd2d3e4&redirect_uri=http%3a%2f%2fwww.greatepanda.top%2fwx_api%2fuser_info&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"
            }]
       }]
}`

func MenuCreate(ctx *gin.Context) {
	err := common.CreateMenu([]byte(MenuStr))
	if err != nil {
		fmt.Sprintf("menu create err: %+v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errNum": common.ErrNumMenuCreateErr,
			"errMsg": common.ErrMsg(common.ErrNumMenuCreateErr),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"errNum": common.ErrNumSuccess,
		"errMsg": common.ErrMsg(common.ErrNumSuccess),
	})
	return
}
