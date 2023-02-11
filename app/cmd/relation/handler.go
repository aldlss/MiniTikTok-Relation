package main

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/pack"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/service"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
	log "github.com/sirupsen/logrus"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// FollowAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowAction(ctx context.Context, req *relation.FollowActionRequest) (resp *relation.FollowActionResponse, err error) {
	if req.Id == 0 || req.ToUserId == 0 {
		err := errno.ParamErr.WriteMsg("Params can't be 0")
		log.Error(err.Error())
		return pack.BuildActionResp(err), nil
	}
	if req.Id == req.ToUserId {
		err := errno.ParamErr.WriteMsg("Can't follow self")
		log.Error(err.Error())
		return pack.BuildActionResp(err), nil
	}

	err = service.NewFollowActionService(ctx).FollowAction(req)
	if err != nil {
		log.Error(err.Error())
		return pack.BuildActionResp(err), nil
	}
	return pack.BuildActionResp(errno.Success), nil
}

func (s *RelationServiceImpl) ListFollow(ctx context.Context, req *relation.ListFollowRequest) (resp *relation.ListFollowResponse, err error) {
	followList, err := service.NewListFollowService(ctx).ListFollow(req)
	if err != nil {
		log.Error(err.Error())
		return pack.BuildListFollowResp(nil, err), nil
	}
	return pack.BuildListFollowResp(followList, errno.Success), nil
}

func (s *RelationServiceImpl) ListFans(ctx context.Context, req *relation.ListFansRequest) (resp *relation.ListFansResponse, err error) {
	fansList, err := service.NewListFansService(ctx).ListFans(req)
	if err != nil {
		log.Error(err.Error())
		return pack.BuildListFansResp(nil, err), nil
	}
	return pack.BuildListFansResp(fansList, errno.Success), nil
}

func (s *RelationServiceImpl) ListFriends(ctx context.Context, req *relation.ListFriendsRequest) (resp *relation.ListFriendsResponse, err error) {
	friendsList, err := service.NewListFriendsService(ctx).ListFriends(req)
	if err != nil {
		log.Error(err.Error())
		return pack.BuildListFriendsResp(nil, err), nil
	}
	return pack.BuildListFriendsResp(friendsList, errno.Success), nil
}
