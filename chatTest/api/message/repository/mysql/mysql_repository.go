package mysql

import (
	model "chatTest/api/message/model"
	"context"
	"github.com/jinzhu/gorm"
	elog "golang.swit.dev/pkg/log/v2"
	sutil "golang.swit.dev/pkg/util"
)

type repository struct {
	db *gorm.DB
}

func (r repository) GetMessageList(ctx context.Context, param model.Message)([]model.Message,error){
var list []model.Message

q :=r.db.Select("chan_no, user_id, msg_cont, msg_time FROM message")
if err := q.Find(&list, "chan_no = ?").Error; err != nil {
	elog.Error(err).Send()
		return nil, err
	}
	return list,nil
}

func (r repository) CreateMessage(ctx context.Context, param model.Message) error{
	sql := `
			INSERT INTO message (msg_no, chan_no, user_id, msg_cont, msg_time)
			VALUES (?, ?, ?, ?, ?);`
	if err := r.db.Exec(sutil.ReplaceQuery(sql),param.MsgNo,param.ChanNo,param.UserId,param.MsgCont,param.MsgTime).Error; err!= nil {
		elog.Error(err).Send()
		return err
	}
	return nil

}

func (r repository) DeleteMessage(ctx context.Context, param model.Message) error{
	sql := `
			DELETE FROM message 
			WHERE chan_no = ? AND msg_no = ? AND user_id = ? ;`
	if err := r.db.Exec(sutil.ReplaceQuery(sql),param.MsgNo,param.ChanNo,param.UserId).Error; err!= nil {
		elog.Error(err).Send()
		return err
	}
	return nil
}

func (r repository) UpdateMessage(ctx context.Context, param model.Message) error{
	sql := `
		UPDATE message SET msg_cont = ? WHERE = chan_no = ? AND msg_no = ? AND user_id = ? `
	if err := r.db.Exec(sutil.ReplaceQuery(sql),param.MsgNo,param.ChanNo,param.UserId,param.MsgCont).Error; err!= nil {
		elog.Error(err).Send()
		return err
	}
	return nil
}