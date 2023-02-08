package main

import (
	"context"
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/api/rpc"
	"github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/auth"
	"github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/relation"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/constants"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/errno"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestRelationApi(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	cli, err := client.NewClient()
	a.NoError(err)

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

	baseUrl := "http://[::1]:19810"
	status, _, err := cli.Get(ctx, nil, baseUrl+"/douyin/relation/aya/list?user_id=111&&token=1061wg")
	a.NoError(err)
	a.Equal(constants.BadRequest, status)

	// 数据库环境与 relation_service_test.go 相同，也照着同样的逻辑写了--
	// aya-11 satori-12 remilia-13 koishi-14
	var wg sync.WaitGroup
	wg.Add(4)
	followAction := func(token string, toId uint32, ActionType relation.FollowActionRequest_FollowActionType) {
		status, body, err := cli.Post(ctx, nil,
			fmt.Sprintf("%v/douyin/relation/action?token=%v&to_user_id=%v&action_type=%d", baseUrl, token, toId, ActionType), nil)
		a.NoError(err)
		a.Equal(constants.OK, status)
		var resp relation.FollowActionResponse
		err = sonic.Unmarshal(body, &resp)
		a.NoError(err)
		a.Zero(resp.StatusCode)
		if ActionType == relation.FollowActionRequest_FOLLOW {
			wg.Done()
		}
	}

	go followAction("aya", 12, relation.FollowActionRequest_FOLLOW)
	go followAction("remilia", 12, relation.FollowActionRequest_FOLLOW)
	go followAction("koishi", 12, relation.FollowActionRequest_FOLLOW)
	go followAction("satori", 14, relation.FollowActionRequest_FOLLOW)
	wg.Wait()

	wg.Add(2)
	go func() {
		status, body, err := cli.Get(ctx, nil,
			baseUrl+"/douyin/relation/follower/list/?token=aya&user_id=12")
		a.NoError(err)
		a.Equal(constants.OK, status)

		var resp relation.ListFansResponse
		a.NoError(sonic.Unmarshal(body, &resp))

		a.Zero(resp.StatusCode)
		a.Equal(3, len(resp.UserList))
		a.False(resp.UserList[2].IsFollow)
		wg.Done()
	}()
	go func() {
		status, body, err := cli.Get(ctx, nil,
			baseUrl+"/douyin/relation/friend/list?user_id=14&token=remilia")
		a.NoError(err)
		a.Equal(constants.OK, status)

		var resp relation.ListFriendsResponse
		a.NoError(sonic.Unmarshal(body, &resp))

		a.Zero(resp.StatusCode)
		a.Equal(1, len(resp.UserList))
		a.True(resp.UserList[0].IsFollow)
		wg.Done()
	}()
	wg.Wait()

	followAction("koishi", 12, relation.FollowActionRequest_UNFOLLOW)

	wg.Add(3)
	go func() {
		status, body, err := cli.Get(ctx, nil,
			baseUrl+"/douyin/relation/follow/list?user_id=14&token=satori")
		a.NoError(err)
		a.Equal(constants.OK, status)

		var resp relation.ListFollowResponse
		a.NoError(sonic.Unmarshal(body, &resp))

		a.Zero(resp.StatusCode)
		a.Zero(len(resp.UserList))
		wg.Done()
	}()
	go func() {
		status, body, err := cli.Post(ctx, nil,
			baseUrl+"/douyin/relation/action/?user_id=11&token=satori&action_type=0", nil)
		a.NoError(err)
		a.Equal(constants.OK, status)

		var resp relation.FollowActionResponse
		a.NoError(sonic.Unmarshal(body, &resp))

		a.EqualValues(resp.StatusCode, errno.ParamErrCode)
		wg.Done()
	}()

	go func() {
		status, body, err := cli.Post(ctx, nil,
			baseUrl+"/douyin/relation/action/?user_id=10&token=satori&action_type=-1", nil)
		a.NoError(err)
		a.Equal(constants.OK, status)

		var resp relation.FollowActionResponse
		a.NoError(sonic.Unmarshal(body, &resp))

		// 期望有 `level=error msg="validating: expr_path=actionType, cause=invalid"` 的 log
		a.EqualValues(resp.StatusCode, errno.UnclassifiedErrCode)
		wg.Done()
	}()
	wg.Wait()
}

func TestMain(m *testing.M) {
	go main()
	time.Sleep(time.Second / 2)
	m.Run()
}
