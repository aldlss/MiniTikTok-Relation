package pack

import (
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/relation/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {
	user := User(&model.User{
		Id:            1,
		Name:          "satori",
		FollowCount:   514,
		FollowerCount: 996,
	}, true)
	assert.NotNil(t, user)
	assert.True(t, user.IsFollow)
	assert.Equal(t, "satori", user.Name)
}

func TestUsers(t *testing.T) {
	users := Users([]*model.User{
		&model.User{
			Id:            1,
			Name:          "Aya",
			FollowCount:   1111,
			FollowerCount: 2222,
		},
		&model.User{
			Id:            5,
			Name:          "Satori",
			FollowCount:   5555,
			FollowerCount: 1111,
		},
		&model.User{
			Id:            9,
			Name:          "Cirno",
			FollowCount:   9,
			FollowerCount: 9,
		},
	}, []*model.User{
		&model.User{
			Id:            5,
			Name:          "Satori",
			FollowCount:   5555,
			FollowerCount: 1111,
		},
	})
	assert.Equal(t, 3, len(users))
	assert.True(t, users[1].IsFollow)
	assert.False(t, users[0].IsFollow)
	assert.Equal(t, "Cirno", users[2].Name)
}
