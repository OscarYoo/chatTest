package model

import "time"

type (
	GetMessageListParam struct {
		ChanNo  int       `json:"chan_no"`
	}
	CreateMessageParam struct {
		MsgNo	int		  `json:"msg_no"`
		ChanNo  int       `json:"chan_no"`
		UserId  string    `json:"user_id"`
		MsgCont string    `json:"msg_cont"`
		MsgTime time.Time `json:"msg_time"`
	}
	UpdateMessageParam struct {
		MsgNo   int       `json:"msg_no"`
		ChanNo  int       `json:"chan_no"`
		UserId  string    `json:"user_id"`
		MsgCont string    `json:"msg_cont"`
	}
	DeleteMessageParam struct {
		MsgNo   int       `json:"msg_no"`
		ChanNo  int       `json:"chan_no"`
		UserId  string    `json:"user_id"`
	}
)