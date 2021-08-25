package model

import "time"

type ChannelListResponse struct {
	Channels []Channel
}
type Channel struct {
	ChanNo   int          `json:"chan_no"`
	UserId   string       `json:"user_id"`
	ChanName string       `json:"chan_name"`
	ChanTime time.Time    `json:"chan_time"`

}

type User struct {
	UserId string      `json:"user_id"`
}