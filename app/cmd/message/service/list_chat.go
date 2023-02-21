package service

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/message/dal/db"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/message/pack"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
)

type ListChatService struct {
	ctx context.Context
}

func NewListChatService(ctx context.Context) *ListChatService {
	return &ListChatService{ctx: ctx}
}

func (s ListChatService) ListChat(req *message.ChatRequest) ([]*message.Message, error) {
	if req.Id == 0 || req.ToUserId == 0 {
		return nil, errno.ParamErr
	}

	dbMessages, err := db.ListFriendMessage(s.ctx, req.Id, req.ToUserId, req.PreMsgTime)
	if err != nil {
		return nil, err
	}

	messages := pack.Messages(dbMessages, req.Id, req.ToUserId)
	return messages, nil
}
