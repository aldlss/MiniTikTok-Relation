package model

type ChatExtend struct {
	Content  string `gorm:"column:content"`
	SenderId int64  `gorm:"column:sender"`
}
