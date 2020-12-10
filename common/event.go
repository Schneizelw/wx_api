package common

type UserBaseEvent struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Event        string
	EventKey     string
}

type UserEventSubscribe struct {
	UserBaseEvent
}

type UserEventUnsubscribe struct {
	UserBaseEvent
}

type UserEventClick struct {
	UserBaseEvent
}
