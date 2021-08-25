package model

import "time"

type (

	CreateChannelParam struct {
		ChanNo   int          `json:"chan_no"`
		UserId   string       `json:"user_id"`
		ChanName string       `json:"chan_name"`
		ChanTime time.Time    `json:"chan_time"`
	}

	GetAllChannelListParam struct {
	}

	GetMyChannelListParam struct {
		UserId   string       `json:"user_id"`
	}
)