package service

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/dal/db"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/pack"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
)

type ListFollowService struct {
	ctx context.Context
}

func NewListFollowService(ctx context.Context) *ListFollowService {
	return &ListFollowService{ctx: ctx}
}

func (s *ListFollowService) ListFollow(req *relation.ListFollowRequest) ([]*relation.User, error) {
	dbUser, err := db.ListRelation(s.ctx, req.UserId, db.FOLLOW)
	if err != nil {
		return nil, err
	}

	followList, err := db.ListRelationWithUserFollow(s.ctx, req.Id, req.UserId, db.FOLLOW)
	if err != nil {
		return nil, err
	}

	return pack.Users(dbUser, followList), nil
}
