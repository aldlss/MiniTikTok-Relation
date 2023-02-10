package main

import (
	"context"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/message/pack"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/message/service"
	"github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/message"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/errno"
	log "github.com/sirupsen/logrus"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// Action implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) Action(ctx context.Context, req *message.ActionRequest) (resp *message.ActionResponse, err error) {
	err = service.NewChatActionService(ctx).ChatActionService(req)
	if err != nil {
		log.Error(err.Error())
		return pack.BuildChatActionResp(err), nil
	}
	return pack.BuildChatActionResp(errno.Success), nil
}

// Chat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) Chat(ctx context.Context, req *message.ChatRequest) (resp *message.ChatResponse, err error) {
	messageList, err := service.NewListChatService(ctx).ListChat(req)
	if err != nil {
		log.Error(err.Error())
		return pack.BuildListChatResp(nil, err), nil
	}
	return pack.BuildListChatResp(messageList, errno.Success), nil
}
