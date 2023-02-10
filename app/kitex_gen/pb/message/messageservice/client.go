// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messageservice

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Action(ctx context.Context, Req *message.ActionRequest, callOptions ...callopt.Option) (r *message.ActionResponse, err error)
	Chat(ctx context.Context, Req *message.ChatRequest, callOptions ...callopt.Option) (r *message.ChatResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMessageServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMessageServiceClient struct {
	*kClient
}

func (p *kMessageServiceClient) Action(ctx context.Context, Req *message.ActionRequest, callOptions ...callopt.Option) (r *message.ActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Action(ctx, Req)
}

func (p *kMessageServiceClient) Chat(ctx context.Context, Req *message.ChatRequest, callOptions ...callopt.Option) (r *message.ChatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Chat(ctx, Req)
}
