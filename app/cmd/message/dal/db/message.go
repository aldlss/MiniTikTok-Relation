package db

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        uint32    `gorm:"primaryKey;column:id"`
	CreatedAt time.Time `gorm:"column:create_time"`
	ChatId    uint64    `gorm:"column:chat_id"`
	Content   string    `gorm:"column:content"`
	SenderId  uint32    `gorm:"column:sender"`
}

func SendFriendMessage(ctx context.Context, fromUserId uint32, toUserId uint32, content string) error {
	from := fromUserId
	if fromUserId > toUserId {
		fromUserId, toUserId = toUserId, fromUserId
	}
	return sendMessage(ctx, uint64(fromUserId)<<32|uint64(toUserId), from, content)
}

func ListFriendMessage(ctx context.Context, fromUserId uint32, toUserId uint32) ([]*Message, error) {
	if fromUserId > toUserId {
		fromUserId, toUserId = toUserId, fromUserId
	}
	return listMessage(ctx, uint64(fromUserId)<<32|uint64(toUserId))
}

func sendMessage(ctx context.Context, chatId uint64, senderId uint32, content string) error {
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

func listMessage(ctx context.Context, chatId uint64) ([]*Message, error) {
	messages := make([]*Message, 0)
	session := PgDb.Session(&gorm.Session{
		SkipHooks: true,
		Context:   ctx,
	})
	res := session.Where("chat_id = ?", chatId).Find(&messages)
	if res.Error != nil {
		return nil, res.Error
	}
	return messages, nil
}
