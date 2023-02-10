package rpc

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/auth"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/auth/authservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
	log "github.com/sirupsen/logrus"
	"time"
)

var authClient authservice.Client

func initAuthRPC() {
	r, err := consul.NewConsulResolverWithConfig(&api.Config{
		Address: "10.233.70.76:14514",
		Scheme:  "http",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	cli, err := authservice.NewClient("minitiktok-auth-grpc",
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "minitiktok-auth-grpc",
		}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*3),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	authClient = cli
}

func AuthRPC(ctx context.Context, req *auth.AuthRequest) (uint32, error) {
	resp, err := authClient.Auth(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != auth.AuthResponse_SUCCESS {
		log.Info("Auth Fail")
		return 0, nil
	}
	return resp.UserId, nil
}
