package handle

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/api/rpc"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/constants"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	log "github.com/sirupsen/logrus"
)

func RelationFollowAction(ctx context.Context, c *app.RequestContext) {
	type followReqWithoutId struct {
		toId       uint32                                        `query:"to_user_id" vd:"$>=0"`
		actionType relation.FollowActionRequest_FollowActionType `query:"action_type" vd:"$>=0"`
	}
	var req followReqWithoutId
	err := c.BindAndValidate(&req)
	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	id, isExist := c.Get("id")
	if !isExist {
		log.Error(errno.NilValueErr)
		SendBaseResponse(c, errno.NilValueErr)
		return
	}

	err = rpc.FollowActionRPC(ctx, &relation.FollowActionRequest{
		Id:         id.(uint32),
		ToUserId:   req.toId,
		ActionType: req.actionType,
	})
	if err != nil {
		log.Error(err.Error())
	}
	SendBaseResponse(c, err)
}

func RelationList(ctx context.Context, c *app.RequestContext) {
	type listRelationWithoutId struct {
		toId uint32 `query:"user_id" vd:"$>=0"`
	}
	var req listRelationWithoutId

	err := c.BindAndValidate(&req)
	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	id, isExist := c.Get("id")
	if !isExist {
		log.Error(errno.NilValueErr)
		SendBaseResponse(c, errno.NilValueErr)
		return
	}

	var resp any
	switch c.Param("relation") {
	case "follow":
		resp, err = rpc.ListFollowRPC(ctx, &relation.ListFollowRequest{
			Id:     id.(uint32),
			UserId: req.toId,
		})
	case "follower":
		resp, err = rpc.ListFansRPC(ctx, &relation.ListFansRequest{
			Id:     id.(uint32),
			UserId: req.toId,
		})
	case "friend":
		resp, err = rpc.ListFriendsRPC(ctx, &relation.ListFriendsRequest{
			Id:     id.(uint32),
			UserId: req.toId,
		})
	default:
		c.String(constants.BadRequest, "Unknown Relation")
		return
	}

	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}
	SendResponse(c, resp)
}
