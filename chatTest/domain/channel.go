package domain

import (
	"chatTest/api/channel/model"
	"context"
	"time"
)

type ChannelRepository interface {
	GetAllChannelList(ctx context.Context, param model.Channel)([]model.Channel, error)
	GetMyChannelList(ctx context.Context, param model.Channel)([]model.Channel, error)
	CreateChannel(ctx context.Context, param model.Channel) error

}

type ChannelService interface {
	GetAllChannelList(ctx context.Context, param model.GetAllChannelListParam)([]model.Channel, error)
	GetMyChannelList(ctx context.Context, param model.GetMyChannelListParam)([]model.Channel, error)
	CreateChannel(ctx context.Context, param model.CreateChannelParam) (model.Channel,error)
}

type Channel struct {
	ChanNo   int          `json:"chan_no"`
	UserId   string       `json:"user_id"`
	ChanName string       `json:"chan_name"`
	ChanTime time.Time    `json:"chan_time"`

}
