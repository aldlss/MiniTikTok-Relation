package rpc

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation/relationservice"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	log "github.com/sirupsen/logrus"
)

var relationClient relationservice.Client

func initRelationRPC() {
	cli, err := relationservice.NewClient("minitiktok-relation-grpc",
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "minitiktok-relation-grpc",
		}),
		client.WithHostPorts("127.0.0.1:19198"),
		//client.WithTransportProtocol(transport.GRPC),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	relationClient = cli
}

func FollowActionRPC(ctx context.Context, req *relation.FollowActionRequest) error {
	resp, err := relationClient.FollowAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

func ListFollowRPC(ctx context.Context, req *relation.ListFollowRequest) (*relation.ListFollowResponse, error) {
	resp, err := relationClient.ListFollow(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func ListFansRPC(ctx context.Context, req *relation.ListFansRequest) (*relation.ListFansResponse, error) {
	resp, err := relationClient.ListFans(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func ListFriendsRPC(ctx context.Context, req *relation.ListFriendsRequest) (*relation.ListFriendsResponse, error) {
	resp, err := relationClient.ListFriends(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}
