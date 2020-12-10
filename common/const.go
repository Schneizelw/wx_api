package common

const (
	WxAppId     = "wxa03e08dedbd2d3e4"
	WxAppSecret = "6e74b2f686e375b99fcc28a8881c5bf9"
)

const (
	WxApiCommonAccessToken       = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	WxApiMenuCreate              = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
	WxApiCommonAccessTokenOauth2 = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	WxApiUserInfoOauth2          = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
)

const (
	WxCheckServerToken = "wuyangqiao"
)

const (
	MsgText  = "text"
	MsgImage = "image"
)

const (
	Event          = "event"
	EventSeparator = "_"
	Subscribe      = "subscribe"
	Unsubscribe    = "unsubscribe"

	EventSubscribe   = Event + EventSeparator + Subscribe
	EventUnsubscribe = Event + EventSeparator + Unsubscribe
)

const (
	Click           = "CLICK"
	ClickV0001      = "click_v0001"
	ClickV0002      = "click_v0002"
	EventClickV0001 = Event + EventSeparator + Click + EventSeparator + ClickV0001
	EventClickV0002 = Event + EventSeparator + Click + EventSeparator + ClickV0002
)
