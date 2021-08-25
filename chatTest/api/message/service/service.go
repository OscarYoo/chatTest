package service

import (
	"chatTest/domain"
	"context"
	"github.com/asaskevich/govalidator"
	"chatTest/api/message/model"
)

type service struct {
	mysql domain.MessageRepository
}

func (s* service)CreateMessage(ctx context.Context, param model.CreateMessageParam)(model.Message,error)  {
	if _, err := govalidator.ValidateStruct(param); err!={
		return model.Message{}, err
	}
	tab := model.Message{
		MsgNo: param.MsgNo,
		ChanNo: param.ChanNo,
		UserId: param.UserId,
		MsgCont: param.MsgCont,
		MsgTime: param.MsgTime,
	}
	err :=s.mysql.CreateMessage(ctx, tab)
	if err != nil {
		return model.Message{}, err
	}

	return tab, nil
}

func (s *service) UpdateMessage(ctx context.Context, param model.UpdateMessageParam) (model.Message, error){
	if _, err := govalidator.ValidateStruct(param); err !=nil {
		return model.Message{},err
	}

	tab :=model.Message{
		MsgNo: param.MsgNo,
		UserId: param.UserId,
		MsgCont: param.MsgCont,
	}
	err := s.mysql.UpdateMessage(ctx, tab)
	if err != nil {
		return model.Message{}, err
	}
	return tab, nil
}

func (s *service) DeleteMessage(ctx context.Context, param model.DeleteMessageParam) error {
	if _,err := govalidator.ValidateStruct(param); err !=nil {
		return err
	}
	err := s.mysql.DeleteMessage(ctx, param.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetMessageList(ctx context.Context, param model.GetMessageListParam)([]model.Message, error)  {
	if _,err := govalidator.ValidateStruct(param); err != nil {
		return nil,err
	}
	res, err := s.mysql.GetMessageList(ctx, param.ChanNo)

	if err != nil {
		return nil, err
	}
	return res, nil
}