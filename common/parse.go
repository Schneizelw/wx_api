package common

import (
	"encoding/xml"
)

func ParseMsgType(data []byte) (string, error) {
	msgType := MsgType{}
	if err := xml.Unmarshal(data, &msgType); err != nil {
		return "", err
	}
	res := msgType.MsgType
	if msgType.Event != "" {
		res += EventSeparator + msgType.Event
	}
	if msgType.EventKey != "" {
		res += EventSeparator + msgType.EventKey
	}
	return res, nil
}

func ParseMsgText(data []byte) (UserMsgText, error) {
	userMsg := UserMsgText{}
	if err := xml.Unmarshal(data, &userMsg); err != nil {
		return userMsg, err
	}
	return userMsg, nil
}

func ParseMsgImage(data []byte) (UserMsgImage, error) {
	userMsg := UserMsgImage{}
	if err := xml.Unmarshal(data, &userMsg); err != nil {
		return userMsg, err
	}
	return userMsg, nil
}

func ParseToXml(respMsg interface{}) (string, error) {
	data, err := xml.MarshalIndent(respMsg, "", "")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ParseEventSubscribe(data []byte) (UserEventSubscribe, error) {
	userEvent := UserEventSubscribe{}
	if err := xml.Unmarshal(data, &userEvent); err != nil {
		return userEvent, err
	}
	return userEvent, nil
}

func ParseEventUnsubscribe(data []byte) (UserEventUnsubscribe, error) {
	userEvent := UserEventUnsubscribe{}
	if err := xml.Unmarshal(data, &userEvent); err != nil {
		return userEvent, err
	}
	return userEvent, nil
}

func ParseEventClick(data []byte) (UserEventClick, error) {
	userEvent := UserEventClick{}
	if err := xml.Unmarshal(data, &userEvent); err != nil {
		return userEvent, err
	}
	return userEvent, nil
}
