package model

import "time"

type MessageListResponse struct {
	Messages []Message
}

type Message struct {
	MsgNo   int       `json:"msg_no"`
	ChanNo  int       `json:"chan_no"`
	UserId  string    `json:"user_id"`
	MsgCont string    `json:"msg_cont"`
	MsgTime time.Time `json:"msg_time"`

}

type User struct {
	UserId string      `json:"user_id"`
}