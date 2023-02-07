package main

import (
	"context"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/relation/pack"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/relation/service"
	relation "github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/relation"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// FollowAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowAction(ctx context.Context, req *relation.FollowActionRequest) (resp *relation.FollowActionResponse, err error) {
	if req.Id == 0 || req.ToUserId == 0 {
		return pack.BuildActionResp(errno.ParamErr.WriteMsg("Params can't be 0")), nil
	}
	if req.Id == req.ToUserId {
		return pack.BuildActionResp(errno.ParamErr.WriteMsg("Can't follow self")), nil
	}

	err = service.NewFollowActionService(ctx).FollowAction(req)
	if err != nil {
		return pack.BuildActionResp(err), err
	}
	return pack.BuildActionResp(errno.Success), nil
}

func (s *RelationServiceImpl) ListFollow(ctx context.Context, req *relation.ListFollowRequest) (resp *relation.ListFollowResponse, err error) {
	followList, err := service.NewListFollowService(ctx).ListFollow(req)
	if err != nil {
		return pack.BuildListFollowResp(nil, err), err
	}
	return pack.BuildListFollowResp(followList, errno.Success), nil
}

func (s *RelationServiceImpl) ListFans(ctx context.Context, req *relation.ListFansRequest) (resp *relation.ListFansResponse, err error) {
	fansList, err := service.NewListFansService(ctx).ListFans(req)
	if err != nil {
		return pack.BuildListFansResp(nil, err), err
	}
	return pack.BuildListFansResp(fansList, errno.Success), nil
}

func (s *RelationServiceImpl) ListFriends(ctx context.Context, req *relation.ListFriendsRequest) (resp *relation.ListFriendsResponse, err error) {
	friendsList, err := service.NewListFriendsService(ctx).ListFriends(req)
	if err != nil {
		return pack.BuildListFriendsResp(nil, err), err
	}
	return pack.BuildListFriendsResp(friendsList, errno.Success), nil
}
