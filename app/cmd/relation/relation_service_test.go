package main

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/dal/db"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/model"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation/relationservice"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"sync"
	"testing"
	"time"
)

func TestRelationService(t *testing.T) {

	cli, err := relationservice.NewClient("minitiktok-relation-grpc",
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "minitiktok-relation-grpc",
		}),
		client.WithHostPorts("127.0.0.1:19198"),
		client.WithTransportProtocol(transport.GRPC),
	)
	assert.Nil(t, err)

	ctx := context.Background()

	var wg sync.WaitGroup
	wg.Add(4)
	followAction := func(id uint32, toId uint32, ActionType relation.FollowActionRequest_FollowActionType) {
		req := &relation.FollowActionRequest{
			Id:         id,
			ToUserId:   toId,
			ActionType: ActionType,
		}
		followResp, err := cli.FollowAction(ctx, req)
		if err != nil {
			log.Error(err.Error())
		}
		assert.Zero(t, followResp.StatusCode)
		if ActionType == relation.FollowActionRequest_FOLLOW {
			wg.Done()
		}
	}

	go followAction(11, 12, relation.FollowActionRequest_FOLLOW)
	go followAction(13, 12, relation.FollowActionRequest_FOLLOW)
	go followAction(14, 12, relation.FollowActionRequest_FOLLOW)
	go followAction(12, 14, relation.FollowActionRequest_FOLLOW)
	wg.Wait()

	wg.Add(2)
	go func() {
		listFansResp, _ := cli.ListFans(ctx, &relation.ListFansRequest{
			Id:     11,
			UserId: 12,
		})
		assert.Zero(t, listFansResp.StatusCode)
		assert.Equal(t, 3, len(listFansResp.UserList))
		assert.False(t, listFansResp.UserList[2].IsFollow)
		wg.Done()
	}()
	go func() {
		listFriResp, _ := cli.ListFriends(ctx, &relation.ListFriendsRequest{
			Id:     13,
			UserId: 14,
		})
		assert.Zero(t, listFriResp.StatusCode)
		assert.Equal(t, 1, len(listFriResp.UserList))
		assert.True(t, listFriResp.UserList[0].IsFollow)
		assert.EqualValues(t, 1, listFriResp.UserList[0].MsgType)
		assert.Equal(t, "I agree with you!", listFriResp.UserList[0].Message)
		wg.Done()
	}()
	wg.Wait()

	followAction(14, 12, relation.FollowActionRequest_UNFOLLOW)

	wg.Add(2)
	go func() {
		listFollowResp, _ := cli.ListFollow(ctx, &relation.ListFollowRequest{
			Id:     12,
			UserId: 14,
		})
		assert.Zero(t, listFollowResp.StatusCode)
		assert.Zero(t, len(listFollowResp.UserList))
		wg.Done()
	}()

	go func() {
		followActionResp, _ := cli.FollowAction(ctx, &relation.FollowActionRequest{
			Id:         11,
			ToUserId:   11,
			ActionType: relation.FollowActionRequest_FOLLOW,
		})
		assert.EqualValues(t, errno.ParamErrCode, followActionResp.StatusCode)
		wg.Done()
	}()
	wg.Wait()
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	err := os.Setenv("TABLE_NAME", "testRelation112_message")
	if err != nil {
		log.Fatal(err.Error())
	}

	//Init()
	go main()
	for db.Driver == nil || db.PgDb == nil {
		time.Sleep(time.Millisecond / 2)
	}

	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: "neo4j",
	})
	_, err = session.Run(ctx, `
		MATCH (a:User)
		WHERE a.name="ayaa" OR a.name="satoria" OR a.name="remiliaa" OR a.name="koishia"
		DETACH DELETE a`,
		map[string]any{})
	if err != nil {
		println(err.Error())
		return
	}
	_, err = session.Run(ctx, `
		CREATE (:User{user_id:11,follower_count:0,follow_count:0,name:"ayaa"}),
		(:User{user_id:12,follower_count:0,follow_count:0,name:"satoria"}),
		(:User{user_id:13,follower_count:0,follow_count:0,name:"remiliaa"}),
		(:User{user_id:14,follower_count:0,follow_count:0,name:"koishia"})`,
		map[string]any{})
	if err != nil {
		log.Error(err)
		return
	}

	err = db.PgDb.Migrator().AutoMigrate(&model.Message{})
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.PgDb.Session(&gorm.Session{
		SkipHooks: true,
	}).Where("1=1").Delete(&model.Message{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.PgDb.Create(&model.Message{
		ChatId:   12<<32 | 13,
		Content:  "Aya, yes!",
		SenderId: 12,
	}).Create(&model.Message{
		ChatId:   12<<32 | 13,
		Content:  "I agree with you!",
		SenderId: 13,
	}).Create(&model.Message{
		ChatId:   13<<32 | 12,
		Content:  "Thank you!",
		SenderId: 13,
	}).Error
	if err != nil {
		log.Fatal()
	}

	m.Run()

	_, err = session.Run(ctx, `
		MATCH (a:User)
		WHERE a.name="ayaa" OR a.name="satoria" OR a.name="remiliaa" OR a.name="koishia"
		DETACH DELETE a`,
		map[string]any{})
	if err != nil {
		log.Error(err.Error())
		return
	}

	err = db.PgDb.Migrator().DropTable(os.Getenv("TABLE_NAME"))
	if err != nil {
		log.Error(err.Error())
	}

	err = session.Close(ctx)
	if err != nil {
		log.Error(err.Error())
	}
}
