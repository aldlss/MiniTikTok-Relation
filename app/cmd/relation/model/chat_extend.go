package model

type ChatExtend struct {
	Content  string `gorm:"column:content"`
	SenderId uint32 `gorm:"column:sender"`
}
