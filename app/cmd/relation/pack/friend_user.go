package pack

import (
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/model"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
)

func FriendUser(user *relation.User, extend *model.ChatExtend, id int64) *relation.FriendUser {
	var msgType int32
	if id == extend.SenderId {
		msgType = 1
	} else {
		msgType = 0
	}
	return &relation.FriendUser{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
		Message:         extend.Content,
		MsgType:         msgType,
	}
}

func FriendUsers(users []*relation.User, extends []*model.ChatExtend, id int64) []*relation.FriendUser {
	friendUsers := make([]*relation.FriendUser, len(users))
	for idx, user := range users {
		friendUsers[idx] = FriendUser(user, extends[idx], id)
	}
	return friendUsers
}
