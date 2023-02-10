package service

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/dal/db"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/pack"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
)

type ListFriendsService struct {
	ctx context.Context
}

func NewListFriendsService(ctx context.Context) *ListFriendsService {
	return &ListFriendsService{ctx: ctx}
}

func (s *ListFriendsService) ListFriends(req *relation.ListFriendsRequest) ([]*relation.User, error) {
	dbUser, err := db.ListRelation(s.ctx, req.UserId, db.FRIENDS)
	if err != nil {
		return nil, err
	}

	friendsList, err := db.ListRelationWithUserFollow(s.ctx, req.Id, req.UserId, db.FRIENDS)
	if err != nil {
		return nil, err
	}

	return pack.Users(dbUser, friendsList), nil
}
