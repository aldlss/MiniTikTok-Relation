package model

import "time"

type Message struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	CreatedAt time.Time `gorm:"column:create_time"`
	ChatId    int64     `gorm:"column:chat_id"`
	Content   string    `gorm:"column:content"`
	SenderId  int64     `gorm:"column:sender"`
}
