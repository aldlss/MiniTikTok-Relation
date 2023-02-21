package pack

import (
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/message/dal/db"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message"
)

func Message(dbMessage *db.Message, fromUserid int64, toUserId int64) *message.Message {
	var toId int64
	if fromUserid == dbMessage.SenderId {
		toId = toUserId
	} else {
		toId = fromUserid
	}
	return &message.Message{
		Id:         dbMessage.ID,
		ToUserId:   toId,
		FromUserId: dbMessage.SenderId,
		Content:    dbMessage.Content,
		CreateTime: uint64(dbMessage.CreatedAt.UnixMilli()),
	}
}

func Messages(dbMessages []*db.Message, fromUserid int64, toUserId int64) []*message.Message {
	messages := make([]*message.Message, len(dbMessages))
	for idx, dbMessage := range dbMessages {
		messages[idx] = Message(dbMessage, fromUserid, toUserId)
	}
	return messages
}
