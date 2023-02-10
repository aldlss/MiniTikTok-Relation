package db

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/model"
	"gorm.io/gorm"
)

func GetFriendExtend(ctx context.Context, fromUserId uint32, toUserId uint32) (*model.ChatExtend, error) {
	if fromUserId > toUserId {
		fromUserId, toUserId = toUserId, fromUserId
	}
	return getChatExtend(ctx, uint64(fromUserId)<<32|uint64(toUserId))
}

func getChatExtend(ctx context.Context, chatId uint64) (*model.ChatExtend, error) {
	chatExtend := &model.ChatExtend{}
	err := PgDb.Session(&gorm.Session{
		SkipHooks: true,
		Context:   ctx,
	}).Where("chat_id = ?", chatId).Order("id DESC").Limit(1).Find(&chatExtend).Error
	if err != nil {
		return nil, err
	}
	return chatExtend, nil
}
