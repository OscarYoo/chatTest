package domain

import (
	"chatTest/api/message/model"
	"context"
	"time"
)

type MessageRepository interface {
	GetMessageList(ctx context.Context, param model.Message)([]model.Message, error)
	CreateMessage(ctx context.Context, param model.Message) error
	UpdateMessage(ctx context.Context, param model.Message) error
	DeleteMessage(ctx context.Context, Userid string) error
}

type MessageService interface {
	GetMessageList(ctx context.Context, param model.GetMessageListParam) ([]model.Message, error)
	CreateMessage(ctx context.Context, param model.CreateMessageParam) (model.Message,error)
	UpdateMessage(ctx context.Context, param model.UpdateMessageParam) (model.Message,error)
	DeleteMessage(ctx context.Context, param model.DeleteMessageParam) error
}

type Message struct {
	MsgNo   int       `json:"msg_no"`
	ChanNo  int       `json:"chan_no"`
	UserId  string    `json:"user_id"`
	MsgCont string    `json:"msg_cont"`
	MsgTime time.Time `json:"msg_time"`

}