package handle

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/api/rpc"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	log "github.com/sirupsen/logrus"
)

func MessageAction(ctx context.Context, c *app.RequestContext) {
	type actionReqWithoutId struct {
		toId       int64  `query:"to_user_id" vd:"$>=0"`
		actionType int32  `query:"action_type" vd:"$==1"`
		content    string `query:"content" vd:"$!=''"`
	}
	var req actionReqWithoutId
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

	err = rpc.SendMessage(ctx, &message.ActionRequest{
		Id:         id.(int64),
		ToUserId:   req.toId,
		ActionType: req.actionType,
		Content:    req.content,
	})
	if err != nil {
		log.Error(err.Error())
	}
	SendBaseResponse(c, err)
}

func MessageChat(ctx context.Context, c *app.RequestContext) {
	type chatReqWithoutId struct {
		toId       int64 `query:"to_user_id" vd:"$>=0"`
		preMsgTime int64 `query:"pre_msg_time" vd:"$>=0"`
	}
	var req chatReqWithoutId
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

	resp, err := rpc.ListChat(ctx, &message.ChatRequest{
		Id:       id.(int64),
		ToUserId: req.toId,
	})
	if err != nil {
		log.Error(err)
		SendBaseResponse(c, err)
		return
	}
	SendResponse(c, resp)
}
