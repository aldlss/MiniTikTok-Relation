package main

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/message/dal/db"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message/messageservice"
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"sync"
	"testing"
	"time"
)

func TestMessageService(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)
	a := assert.New(t)

	cli, err := messageservice.NewClient("minitiktok-message-grpc",
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "minitiktok-message-grpc",
		}),
		client.WithHostPorts("[::1]:15677"),
		client.WithTransportProtocol(transport.GRPC),
	)
	r.NoError(err)

	send := func(id uint32, toId uint32, content string) {
		resp, err := cli.Action(ctx, &message.ActionRequest{
			Id:         id,
			ToUserId:   toId,
			ActionType: 1,
			Content:    content,
		})
		r.NoError(err)
		r.Zero(resp.StatusCode)
	}
	send(1, 2, "hello?")
	send(2, 1, "Bye~")
	send(1, 2, "Do u like Satori sama?")
	send(2, 1, "I think I prefer Aya.")

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		resp, err := cli.Chat(ctx, &message.ChatRequest{
			Id:       1,
			ToUserId: 2,
		})
		a.NoError(err)
		a.Zero(resp.StatusCode)
		a.EqualValues(4, len(resp.MessageList))
		addend := 0
		for idx, msg := range resp.MessageList {
			a.EqualValues(1+addend, msg.FromUserId)
			addend ^= 1
			log.Infof("%d %s %d->%d:%s", idx, msg.CreateTime, msg.FromUserId, msg.ToUserId, msg.Content)
		}
		wg.Done()
	}()

	go func() {
		resp, err := cli.Action(ctx, &message.ActionRequest{
			Id:         2,
			ToUserId:   2,
			ActionType: 5,
			Content:    "test",
		})
		a.NoError(err)
		a.EqualValues(resp.StatusCode, errno.ParamErrCode)
		wg.Done()
	}()

	go func() {
		resp, err := cli.Chat(ctx, &message.ChatRequest{
			Id:       5,
			ToUserId: 0,
		})
		a.NoError(err)
		a.EqualValues(resp.StatusCode, errno.ParamErrCode)
		wg.Done()
	}()

	wg.Wait()
}

func TestMain(m *testing.M) {
	err := db.PgDb.Session(&gorm.Session{
		SkipHooks: true,
	}).Where("1=1").Delete(&db.Message{}).Error
	if err != nil {
		log.Fatal(err.Error())
	}

	go main()
	time.Sleep(time.Second / 2)
	m.Run()

	err = db.PgDb.Migrator().DropTable(&db.Message{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
