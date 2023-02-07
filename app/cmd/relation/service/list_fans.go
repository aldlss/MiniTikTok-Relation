package service

import (
	"context"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/relation/dal/db"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/relation/pack"
	"github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/relation"
)

type ListFansService struct {
	ctx context.Context
}

func NewListFansService(ctx context.Context) *ListFansService {
	return &ListFansService{ctx: ctx}
}

func (s *ListFansService) ListFans(req *relation.ListFansRequest) ([]*relation.User, error) {
	dbUser, err := db.ListRelation(s.ctx, req.UserId, db.FANS)
	if err != nil {
		return nil, err
	}

	fansList, err := db.ListRelationWithUserFollow(s.ctx, req.Id, req.UserId, db.FANS)
	if err != nil {
		return nil, err
	}

	return pack.Users(dbUser, fansList), nil
}
