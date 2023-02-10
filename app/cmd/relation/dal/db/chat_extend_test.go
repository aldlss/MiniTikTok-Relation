package db

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/model"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"sync"
	"testing"
)

type testChatExtend struct {
	suite.Suite
}

func (s *testChatExtend) TestGetFriendExtend() {
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		chatExtend, err := GetFriendExtend(ctx, 5, 0)
		s.NoError(err)
		s.EqualValues("U say right", chatExtend.Content)
		s.EqualValues(5, chatExtend.SenderId)
		wg.Done()
	}()
	go func() {
		chatExtend, err := GetFriendExtend(ctx, 1, 5)
		s.NoError(err)
		s.Empty(chatExtend.Content)
		s.Empty(chatExtend.SenderId)
		wg.Done()
	}()
	wg.Wait()
}

func (s *testChatExtend) SetupSuite() {
	err := os.Setenv("TABLE_NAME", "testRelation111_message")
	if err != nil {
		log.Fatal(err.Error())
	}
	initPgsql()

	err = PgDb.Migrator().AutoMigrate(&model.Message{})
	if err != nil {
		log.Fatal(err.Error())
	}
	err = PgDb.Session(&gorm.Session{
		SkipHooks: true,
	}).Where("1=1").Delete(&model.Message{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err = PgDb.Create(&model.Message{
		ChatId:   5,
		Content:  "Satori, yes!",
		SenderId: 0,
	}).Create(&model.Message{
		ChatId:   5,
		Content:  "U say right",
		SenderId: 5,
	}).Create(&model.Message{
		ChatId:   3,
		Content:  "I am always right",
		SenderId: 5,
	}).Error
	if err != nil {
		log.Fatal(err.Error())
	}

}

func (s *testChatExtend) TearDownSuite() {
	err := PgDb.Migrator().DropTable(os.Getenv("TABLE_NAME"))
	if err != nil {
		log.Error(err.Error())
	}
}

func TestGetFriendExtend(t *testing.T) {
	suite.Run(t, new(testChatExtend))
}
