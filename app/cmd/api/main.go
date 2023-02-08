package main

import (
	"context"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/api/handle"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/api/middleware"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/api/rpc"
	"github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/auth"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/constants"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	rpc.Init()
}

func main() {
	r := server.Default(
		server.WithHostPorts("[::]:19810"),
		server.WithHandleMethodNotAllowed(true),
	)
	Init()
	tiktokGroup := r.Group("/douyin", middleware.Auth)

	relationGroup := tiktokGroup.Group("/relation")
	relationGroup.POST("/action", handle.RelationFollowAction)
	relationGroup.GET("/:relation/list", handle.RelationList)

	patches := gomonkey.ApplyFunc(rpc.AuthRPC, func(ctx context.Context, req *auth.AuthRequest) (uint32, error) {
		switch req.Token {
		case "aya":
			return 11, nil
		case "satori":
			return 12, nil
		case "remilia":
			return 13, nil
		case "koishi":
			return 14, nil
		default:
			return 0, nil
		}
	})
	defer patches.Reset()

	r.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		ctx.Status(constants.NotFound)
	})
	r.NoMethod(func(c context.Context, ctx *app.RequestContext) {
		ctx.Status(constants.MethodNotAllowed)
	})

	r.Spin()
}
