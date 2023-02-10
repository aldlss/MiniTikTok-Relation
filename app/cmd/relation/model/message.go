package model

import "time"

type Message struct {
	ID        uint32    `gorm:"primaryKey;column:id"`
	CreatedAt time.Time `gorm:"column:create_time"`
	ChatId    uint64    `gorm:"column:chat_id"`
	Content   string    `gorm:"column:content"`
	SenderId  uint32    `gorm:"column:sender"`
}
