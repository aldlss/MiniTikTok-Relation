package pack

import (
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/pack"
)

func BuildChatActionResp(err error) *message.ActionResponse {
	baseResp := pack.BuildBaseResp(err)
	return chatActionResp(baseResp)
}

func chatActionResp(baseResp *pack.BaseResp) *message.ActionResponse {
	return &message.ActionResponse{
		StatusCode: baseResp.StatusCode,
		StatusMsg:  baseResp.StatusMsg,
	}
}

func BuildListChatResp(messages []*message.Message, err error) *message.ChatResponse {
	baseResp := pack.BuildBaseResp(err)
	return listChatResp(baseResp, messages)
}

func listChatResp(baseResp *pack.BaseResp, messages []*message.Message) *message.ChatResponse {
	return &message.ChatResponse{
		StatusCode:  baseResp.StatusCode,
		StatusMsg:   baseResp.StatusMsg,
		MessageList: messages,
	}
}
