package service

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/dal/db"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
	log "github.com/sirupsen/logrus"
)

type FollowActionService struct {
	ctx context.Context
}

func NewFollowActionService(ctx context.Context) *FollowActionService {
	return &FollowActionService{ctx: ctx}
}

func (s *FollowActionService) FollowAction(req *relation.FollowActionRequest) error {
	switch req.ActionType {
	case relation.FollowActionRequest_FOLLOW:
		return db.Follow(s.ctx, req.Id, req.ToUserId)
	case relation.FollowActionRequest_UNFOLLOW:
		return db.UnFollow(s.ctx, req.Id, req.ToUserId)
	default:
		log.Error("FollowAction:不支持的操作类型")
		return errno.ParamErr
	}
}
