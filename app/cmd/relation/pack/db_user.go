package pack

import (
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/relation/model"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func DbUser(record *neo4j.Record) *model.User {
	m := record.Values[0].(neo4j.Node).Props

	id, isExist1 := m["user_id"]
	name, isExist2 := m["name"]
	followCount, isExist3 := m["follow_count"]
	followerCount, isExist4 := m["follower_count"]
	avatar, isExist5 := m["avatar"]
	if !isExist5 {
		avatar = ""
	}
	backgroundImage, isExist6 := m["background_image"]
	if !isExist6 {
		backgroundImage = ""
	}

	signature, isExist7 := m["signature"]
	if !isExist7 {
		signature = ""
	}
	totalFavorited, isExist8 := m["total_favorited"]
	if !isExist8 {
		totalFavorited = int64(0)
	}
	workCount, isExist9 := m["work_count"]
	if !isExist9 {
		workCount = int64(0)
	}
	favoriteCount, isExist10 := m["favorite_count"]
	if !isExist10 {
		favoriteCount = int64(0)
	}

	if !(isExist1 && isExist2 && isExist3 && isExist4) {
		return nil
	}

	return &model.User{
		Id:              id.(int64),
		Name:            name.(string),
		FollowCount:     followCount.(int64),
		FollowerCount:   followerCount.(int64),
		Avatar:          avatar.(string),
		BackgroundImage: backgroundImage.(string),
		Signature:       signature.(string),
		TotalFavorited:  totalFavorited.(int64),
		WorkCount:       workCount.(int64),
		FavoriteCount:   favoriteCount.(int64),
	}
}

func DbUsers(records []*neo4j.Record) []*model.User {
	users := make([]*model.User, len(records))
	for idx, record := range records {
		if user := DbUser(record); user != nil {
			users[idx] = user
		}
	}
	return users
}
