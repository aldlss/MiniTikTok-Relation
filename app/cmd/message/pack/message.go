package pack

import (
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/message/dal/db"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message"
)

func Message(dbMessage *db.Message, fromUserid uint32, toUserId uint32) *message.Message {
	var toId uint32
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
		CreateTime: dbMessage.CreatedAt.String(),
	}
}

func Messages(dbMessages []*db.Message, fromUserid uint32, toUserId uint32) []*message.Message {
	messages := make([]*message.Message, len(dbMessages))
	for idx, dbMessage := range dbMessages {
		messages[idx] = Message(dbMessage, fromUserid, toUserId)
	}
	return messages
}
