package service

import (
	"chatTest/domain"
	"context"
	"github.com/asaskevich/govalidator"
	"chatTest/api/channel/model"
)

type service struct {
	mysql domain.ChannelRepository
}

func (s* service)CreateChannel(ctx context.Context, param model.CreateChannelParam)(model.Channel,error)  {
	if _, err := govalidator.ValidateStruct(param); err!= nil{
		return model.Channel{}, err
	}
	tab := model.Channel{
		ChanNo: param.ChanNo,
		UserId: param.UserId,
		ChanName: param.ChanName,
		ChanTime: param.ChanTime,
	}
	err :=s.mysql.CreateChannel(ctx, tab)
	if err != nil {
		return model.Channel{}, err
	}

	return tab, nil
}


func (s *service) GetAllChannelList(ctx context.Context, param model.GetAllChannelListParam)([]model.Channel, error)  {
	if _,err := govalidator.ValidateStruct(param); err != nil {
		return nil,err
	}
	res, err := s.mysql.GetAllChannelList(ctx,?)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) GetMyChannelList(ctx context.Context, param model.GetMyChannelListParam)([]model.Channel, error)  {
	if _,err := govalidator.ValidateStruct(param); err != nil {
		return nil,err
	}
	res, err := s.mysql.GetMyChannelList(ctx, param.UserId)

	if err != nil {
		return nil, err
	}
	return res, nil
}