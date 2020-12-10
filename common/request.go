package common

import (
	"fmt"
	"github.com/guonaihong/gout"
)

type WxBaseResp struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WxRespAccessToken struct {
	WxBaseResp
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
}

type WxRespOauth2AccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

type WxRespOauth2UserInfo struct {
	OpenID     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        string `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	HeadImgUrl string `json:"headimgurl"`
	Privilege  string `json:"privilege"`
	Unionid    string `json:"unionid"`
}

func GetAccessToken() (string, error) {
	url := fmt.Sprintf(WxApiCommonAccessToken, WxAppId, WxAppSecret)
	resp := WxRespAccessToken{}
	err := gout.GET(url).BindJSON(&resp).Do()
	if err != nil {
		fmt.Printf("access token err: %s \n", err)
		return "", err
	}
	fmt.Printf("resp: %+v\n", resp)
	return resp.AccessToken, nil
}

func CreateMenu(menu []byte) error {
	token, err := GetAccessToken()
	if err != nil {
		fmt.Sprintf("get access token err: %+v", err.Error())
		return err
	}
	url := fmt.Sprintf(WxApiMenuCreate, token)
	resp := WxBaseResp{}
	err = gout.POST(url).SetHeader(gout.H{"Content-Type": "application/json; encoding=utf-8"}).
		SetBody(menu).BindJSON(&resp).Do()
	fmt.Printf("menu create resp: %+v", resp)
	if err != nil {
		fmt.Printf("menu create err: %s, resp: %+v \n", err, resp)
		return err
	}
	return nil
}

func GetAccessTokenOauth2(code string) (WxRespOauth2AccessToken, error) {
	url := fmt.Sprintf(WxApiCommonAccessTokenOauth2, WxAppId, WxAppSecret, code)
	resp := WxRespOauth2AccessToken{}
	err := gout.GET(url).BindJSON(&resp).Do()
	if err != nil {
		fmt.Printf("get access token err: %s \n", err)
		return resp, err
	}
	fmt.Printf("resp: %+v\n", resp)
	return resp, nil
}

func GetUserInfoOauth2(accessToken, openID string) (WxRespOauth2UserInfo, error) {
	url := fmt.Sprintf(WxApiUserInfoOauth2, accessToken, openID)
	resp := WxRespOauth2UserInfo{}
	err := gout.GET(url).BindJSON(&resp).Do()
	if err != nil {
		fmt.Printf("get access token err: %s \n", err)
		return resp, err
	}
	fmt.Printf("resp: %+v\n", resp)
	return resp, nil
}
