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
	if !(isExist1 && isExist2 && isExist3 && isExist4) {
		return nil
	}

	return &model.User{
		Id:            uint32(id.(int64)),
		Name:          name.(string),
		FollowCount:   uint32(followCount.(int64)),
		FollowerCount: uint32(followerCount.(int64)),
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
