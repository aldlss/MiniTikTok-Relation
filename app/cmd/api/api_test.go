package main

import (
	"context"
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/api/rpc"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/auth"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/constants"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	log "github.com/sirupsen/logrus"
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
	wg.Add(5)
	sendMessage := func(token string, toId uint32, content string) {
		status, body, err := cli.Post(ctx, nil,
			fmt.Sprintf("%s/douyin/message/action/?token=%s&to_user_id=%d&content=%s&action_type=1",
				baseUrl, token, toId, content), nil)
		a.NoError(err)
		a.Equal(constants.OK, status)

		var resp message.ActionResponse
		a.NoError(sonic.Unmarshal(body, &resp))
		a.Zero(resp.StatusCode)
	}
	contentList := []string{"姐姐~~~", "どうした，恋？", "☑Θ〓∧😭Щㄱeə╋⑨", "Koishi, are you OK?"}
	go func() {
		sendMessage("koishi", 12, contentList[0])
		sendMessage("satori", 14, contentList[1])
		sendMessage("koishi", 12, contentList[2])
		sendMessage("aya", 12, "Ayayayayayayaya")
		sendMessage("satori", 14, contentList[3])
		wg.Done()
	}()

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

	wg.Add(4)
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
	go func() {
		status, body, err := cli.Get(ctx, nil,
			baseUrl+"/douyin/relation/friend/list?user_id=14&token=koishi")
		a.NoError(err)
		a.Equal(constants.OK, status)

		var resp relation.ListFriendsResponse
		a.NoError(sonic.Unmarshal(body, &resp))

		user := resp.UserList[0]
		a.Zero(resp.StatusCode)
		a.Equal(1, len(resp.UserList))
		a.True(user.IsFollow)
		a.EqualValues(user.Name, "satoria")
		a.EqualValues(user.MsgType, 0)
		a.EqualValues(user.Message, "Koishi, are you OK?")
		a.EqualValues(user.FollowerCount, 3)
		a.EqualValues(user.FollowCount, 1)
		wg.Done()
	}()
	go func() {
		status, body, err := cli.Get(ctx, nil,
			baseUrl+"/douyin/message/chat?to_user_id=14&token=satori")
		a.NoError(err)
		a.Equal(constants.OK, status)

		var resp message.ChatResponse
		a.NoError(sonic.Unmarshal(body, &resp))

		a.Zero(resp.StatusCode)
		addend := 1
		for idx, msg := range resp.MessageList {
			log.Infof("%d %d %d->%d: %s", msg.Id, msg.CreateTime, msg.FromUserId, msg.ToUserId, msg.Content)
			a.EqualValues(13+addend, msg.FromUserId)
			a.EqualValues(13-addend, msg.ToUserId)
			addend *= -1
			a.EqualValues(contentList[idx], msg.Content)
			if idx == 3 {
				break
			}
		}
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

		// 期望有 `level=error msg="ErrCode:1001, ErrMsg:Params can't be 0"` 的 log
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
	time.Sleep(time.Second)
	m.Run()
}
