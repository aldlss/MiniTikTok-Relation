package service

import (
	"context"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/message/dal/db"
	"github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/message"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/errno"
)

type ChatActionService struct {
	ctx context.Context
}

func NewChatActionService(ctx context.Context) *ChatActionService {
	return &ChatActionService{ctx: ctx}
}

func (s ChatActionService) ChatActionService(req *message.ActionRequest) error {
	if req.ActionType != 1 || req.Id == 0 || req.ToUserId == 0 {
		return errno.ParamErr
	}
	if req.Content == "" {
		return errno.ParamErr.WriteMsg("content can't be empty")
	}

	err := db.SendFriendMessage(s.ctx, req.Id, req.ToUserId, req.Content)
	if err != nil {
		return err
	}
	return nil
}
