package db

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	CreatedAt time.Time `gorm:"column:create_time"`
	ChatId    int64     `gorm:"column:chat_id"`
	Content   string    `gorm:"column:content"`
	SenderId  int64     `gorm:"column:sender"`
}

func SendFriendMessage(ctx context.Context, fromUserId int64, toUserId int64, content string) error {
	from := fromUserId
	if fromUserId > toUserId {
		fromUserId, toUserId = toUserId, fromUserId
	}
	return sendMessage(ctx, (fromUserId<<32)^toUserId, from, content)
}

func ListFriendMessage(ctx context.Context, fromUserId int64, toUserId int64, preMsgTime int64) ([]*Message, error) {
	if fromUserId > toUserId {
		fromUserId, toUserId = toUserId, fromUserId
	}
	return listMessage(ctx, (fromUserId<<32)^toUserId, preMsgTime)
}

func sendMessage(ctx context.Context, chatId int64, senderId int64, content string) error {
	session := PgDb.Session(&gorm.Session{
		SkipHooks: true,
		Context:   ctx,
	})
	res := session.Create(&Message{
		ChatId:   chatId,
		Content:  content,
		SenderId: senderId,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func listMessage(ctx context.Context, chatId int64, preMsgTime int64) ([]*Message, error) {
	messages := make([]*Message, 0)
	session := PgDb.Session(&gorm.Session{
		SkipHooks: true,
		Context:   ctx,
	})
	res := session.Where("chat_id = ? AND create_time >= ?", chatId, time.Unix(preMsgTime, 0)).Find(&messages)
	if res.Error != nil {
		return nil, res.Error
	}
	return messages, nil
}
