package service

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/dal/db"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/model"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/pack"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
)

type ListFriendsService struct {
	ctx context.Context
}

func NewListFriendsService(ctx context.Context) *ListFriendsService {
	return &ListFriendsService{ctx: ctx}
}

func (s *ListFriendsService) ListFriends(req *relation.ListFriendsRequest) ([]*relation.FriendUser, error) {
	dbUsers, err := db.ListRelation(s.ctx, req.UserId, db.FRIENDS)
	if err != nil {
		return nil, err
	}

	c := make(chan *model.ChatExtend, 5)
	var cancel func()
	s.ctx, cancel = context.WithCancel(s.ctx)
	errChan := make(chan error)

	go func() {
		for _, dbUser := range dbUsers {
			chatExtend, err := db.GetFriendExtend(s.ctx, req.Id, dbUser.Id)
			if err != nil {
				errChan <- err
				return
			}
			select {
			case <-s.ctx.Done():
				return
			case c <- chatExtend:
			}
		}
	}()

	friendsList, err := db.ListRelationWithUserFollow(s.ctx, req.Id, req.UserId, db.FRIENDS)
	if err != nil {
		cancel()
		return nil, err
	}
	users := pack.Users(dbUsers, friendsList)

	friendUsers := make([]*relation.FriendUser, len(dbUsers))
	for idx, user := range users {
		select {
		case err = <-errChan:
			cancel()
			return nil, err
		case chatExtend := <-c:
			friendUsers[idx] = pack.FriendUser(user, chatExtend, req.Id)
		}
	}
	cancel()
	return friendUsers, nil
}
