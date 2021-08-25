package mysql

import (
	model "chatTest/api/channel/model"
	"context"
	"github.com/jinzhu/gorm"
	elog "golang.swit.dev/pkg/log/v2"
	sutil "golang.swit.dev/pkg/util"
)

type repository struct {
	db *gorm.DB
}

func (r repository)CreateChannel(ctx context.Context, param model.Channel) error{
	sql := `
			INSERT INTO channel (chan_no, user_id, chan_name, chan_time)
			VALUES (?, ?, ?, ?);`
	if err := r.db.Exec(sutil.ReplaceQuery(sql),param.ChanNo,param.UserId,param.ChanName,param.ChanTime).Error; err!= nil {
		elog.Error(err).Send()
		return err
	}
	return nil
}

func (r repository) GetAllChannelList(ctx context.Context, param model.Channel)([]model.Channel,error){
	var list []model.Channel

	q :=r.db.Select("chan_no, user_id, chan_name, chan_time FROM channel")
	if err := q.Find(&list).Error; err != nil {
		elog.Error(err).Send()
		return nil, err
	}
	return list,nil

}

func (r repository) GetMyChannelList(ctx context.Context, param model.Channel)([]model.Channel,error){
	var list []model.Channel

	q :=r.db.Select("chan_no, user_id, chan_name, chan_time FROM channel")
	if err := q.Find(&list,"user_id = ?", "?").Error; err != nil {
		elog.Error(err).Send()
		return nil, err
	}
	return list,nil
}
