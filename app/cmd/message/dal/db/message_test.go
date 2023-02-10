package db

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"os"
	"testing"
)

func testSendFriendMessage(t *testing.T) {
	r := require.New(t)
	ctx := context.Background()
	r.NoError(SendFriendMessage(ctx, 0, 0, "satori kawaii!"))
	r.NoError(SendFriendMessage(ctx, 1, 0, "wocao"))
	r.NoError(SendFriendMessage(ctx, 1, 0, ""))
}

func testListFriendMessage(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	res, err := ListFriendMessage(ctx, 0, 0)
	a.NoError(err)
	a.EqualValues(1, len(res))
	a.EqualValues(0, res[0].ChatId)

	res, err = ListFriendMessage(ctx, 1, 0)
	a.NoError(err)
	a.EqualValues(2, len(res))
	a.EqualValues(1, res[0].SenderId)

	log.Info(res[1])
}

func TestArrange(t *testing.T) {
	t.Run("TestSendFriendMessage", testSendFriendMessage)
	t.Run("TestListFriendMessage", testListFriendMessage)
}

func TestMain(m *testing.M) {
	// 使用一个没有应该没有人用的表 test 测试
	err := os.Setenv("TABLE_PREFIX", "test111_")
	if err != nil {
		log.Fatal(err.Error())
	}
	Init()

	err = PgDb.Session(&gorm.Session{
		SkipHooks: true,
	}).Where("1=1").Delete(&Message{}).Error
	if err != nil {
		log.Fatal(err)
	}

	m.Run()

	err = PgDb.Migrator().DropTable(&Message{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
