package pack

import (
	"errors"
	"github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/relation"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/errno"
)

type BaseResp struct {
	StatusCode int32
	StatusMsg  string
}

func buildBaseResp(err error) *BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	e = errno.ServiceErr.WriteMsg(err.Error())
	return baseResp(e)
}

func baseResp(err errno.ErrNo) *BaseResp {
	return &BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}

func BuildActionResp(err error) *relation.FollowActionResponse {
	baseResp := buildBaseResp(err)
	return actionResp(baseResp)
}

func actionResp(baseResp *BaseResp) *relation.FollowActionResponse {
	return &relation.FollowActionResponse{
		StatusCode: baseResp.StatusCode,
		StatusMsg:  baseResp.StatusMsg,
	}
}

func BuildListFollowResp(followList []*relation.User, err error) *relation.ListFollowResponse {
	baseResp := buildBaseResp(err)
	return listFollowResp(baseResp, followList)
}

func listFollowResp(baseResp *BaseResp, followList []*relation.User) *relation.ListFollowResponse {
	return &relation.ListFollowResponse{
		StatusCode: baseResp.StatusCode,
		StatusMsg:  baseResp.StatusMsg,
		UserList:   followList,
	}
}

func BuildListFansResp(fansList []*relation.User, err error) *relation.ListFansResponse {
	baseResp := buildBaseResp(err)
	return listFansResp(baseResp, fansList)
}

func listFansResp(baseResp *BaseResp, fansList []*relation.User) *relation.ListFansResponse {
	return &relation.ListFansResponse{
		StatusCode: baseResp.StatusCode,
		StatusMsg:  baseResp.StatusMsg,
		UserList:   fansList,
	}
}

func BuildListFriendsResp(friendsList []*relation.User, err error) *relation.ListFriendsResponse {
	baseResp := buildBaseResp(err)
	return listFriendsResp(baseResp, friendsList)
}

func listFriendsResp(baseResp *BaseResp, friendsList []*relation.User) *relation.ListFriendsResponse {
	return &relation.ListFriendsResponse{
		StatusCode: baseResp.StatusCode,
		StatusMsg:  baseResp.StatusMsg,
		UserList:   friendsList,
	}
}
