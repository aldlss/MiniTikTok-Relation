package rpc

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message/messageservice"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	log "github.com/sirupsen/logrus"
)

var messageClient messageservice.Client

func initMessageRPC() {
	cli, err := messageservice.NewClient("minitiktok-message-grpc",
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "minitiktok-message-grpc",
		}),
		client.WithHostPorts("127.0.0.1:15677"),
		//client.WithTransportProtocol(transport.GRPC),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	messageClient = cli
}

func SendMessage(ctx context.Context, req *message.ActionRequest) error {
	resp, err := messageClient.Action(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

func ListChat(ctx context.Context, req *message.ChatRequest) (*message.ChatResponse, error) {
	resp, err := messageClient.Chat(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}
