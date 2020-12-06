package common

type WxBaseResp struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WxRespAccessToken struct {
	WxBaseResp
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
}
