package common

type UserBaseMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	MsgId        string
}

type UserMsgText struct {
	UserBaseMsg
	Content string
}

type UserMsgImage struct {
	UserBaseMsg
	PicUrl  string
	MediaId string
}

type RespMsgBase struct {
	ToUserName   string `xml:"xml>ToUserName"` // >表示：<xml>这里加入ToUserName标签</xml>
	FromUserName string `xml:"xml>FromUserName"`
	CreateTime   string `xml:"xml>CreateTime"`
	MsgType      string `xml:"xml>MsgType"`
}

type RespMsgText struct {
	RespMsgBase
	Content string `xml:"xml>Content"`
}

type RespMsgImage struct {
	RespMsgBase
	MediaId string `xml:"xml>Image>MediaId"`
}

type MsgType struct {
	MsgType  string
	Event    string
	EventKey string
}
