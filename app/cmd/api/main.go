package main

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/api/handle"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/api/middleware"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/api/rpc"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/constants"
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

	messageGroup := tiktokGroup.Group("/message")
	messageGroup.POST("/action", handle.MessageAction)
	messageGroup.GET("/chat", handle.MessageChat)

	r.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		ctx.Status(constants.NotFound)
	})
	r.NoMethod(func(c context.Context, ctx *app.RequestContext) {
		ctx.Status(constants.MethodNotAllowed)
	})

	r.Spin()
}
