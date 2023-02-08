package middleware

import (
	"context"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/api/rpc"
	"github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/auth"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/constants"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	log "github.com/sirupsen/logrus"
)

func Auth(c context.Context, ctx *app.RequestContext) {
	token := ctx.Query("token")
	if token == "" {
		err := ctx.AbortWithError(constants.BadRequest, errno.NilValueErr)
		log.Error(err)
		return
	}

	id, err := rpc.AuthRPC(c, &auth.AuthRequest{Token: token})
	if err != nil {
		err = ctx.AbortWithError(constants.InternalServerError, err)
		log.Error(err)
	}

	ctx.Set("id", id)
}
