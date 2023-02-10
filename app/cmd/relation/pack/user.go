package pack

import (
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/model"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/relation"
)

func User(dbUser *model.User, isFollow bool) *relation.User {
	return &relation.User{
		Id:            dbUser.Id,
		Name:          dbUser.Name,
		FollowCount:   dbUser.FollowCount,
		FollowerCount: dbUser.FollowerCount,
		IsFollow:      isFollow,
	}
}

// Users
// 以 dbUsers 构建 relation.User ，并把同时存在于 followList 里的 user 视为已关注
func Users(dbUsers []*model.User, followList []*model.User) []*relation.User {
	followSet := make(map[uint32]struct{}, len(followList))
	for _, dbUser := range followList {
		followSet[dbUser.Id] = struct{}{}
	}
	users := make([]*relation.User, len(dbUsers))
	for idx, dbUser := range dbUsers {
		_, isFollow := followSet[dbUser.Id]
		users[idx] = &relation.User{
			Id:            dbUser.Id,
			Name:          dbUser.Name,
			FollowCount:   dbUser.FollowCount,
			FollowerCount: dbUser.FollowerCount,
			IsFollow:      isFollow,
		}
	}
	return users
}
