package pack

import (
	"github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/relation"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/pack"
)

func BuildActionResp(err error) *relation.FollowActionResponse {
	baseResp := pack.BuildBaseResp(err)
	return actionResp(baseResp)
}

func actionResp(baseResp *pack.BaseResp) *relation.FollowActionResponse {
	return &relation.FollowActionResponse{
		StatusCode: baseResp.StatusCode,
		StatusMsg:  baseResp.StatusMsg,
	}
}

func BuildListFollowResp(followList []*relation.User, err error) *relation.ListFollowResponse {
	baseResp := pack.BuildBaseResp(err)
	return listFollowResp(baseResp, followList)
}

func listFollowResp(baseResp *pack.BaseResp, followList []*relation.User) *relation.ListFollowResponse {
	return &relation.ListFollowResponse{
		StatusCode: baseResp.StatusCode,
		StatusMsg:  baseResp.StatusMsg,
		UserList:   followList,
	}
}

func BuildListFansResp(fansList []*relation.User, err error) *relation.ListFansResponse {
	baseResp := pack.BuildBaseResp(err)
	return listFansResp(baseResp, fansList)
}

func listFansResp(baseResp *pack.BaseResp, fansList []*relation.User) *relation.ListFansResponse {
	return &relation.ListFansResponse{
		StatusCode: baseResp.StatusCode,
		StatusMsg:  baseResp.StatusMsg,
		UserList:   fansList,
	}
}

func BuildListFriendsResp(friendsList []*relation.User, err error) *relation.ListFriendsResponse {
	baseResp := pack.BuildBaseResp(err)
	return listFriendsResp(baseResp, friendsList)
}

func listFriendsResp(baseResp *pack.BaseResp, friendsList []*relation.User) *relation.ListFriendsResponse {
	return &relation.ListFriendsResponse{
		StatusCode: baseResp.StatusCode,
		StatusMsg:  baseResp.StatusMsg,
		UserList:   friendsList,
	}
}
