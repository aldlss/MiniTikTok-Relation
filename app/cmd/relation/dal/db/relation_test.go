package db

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

type testRelation struct {
	suite.Suite
	session neo4j.SessionWithContext
	ctx     context.Context
}

func (s *testRelation) testIsFollow() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		isFollow, err := IsFollow(s.ctx, 1, 2)
		s.NoError(err)
		s.False(isFollow)
		wg.Done()
	}()
	go func() {
		isFollow, err := IsFollow(s.ctx, 5, 6)
		s.NoError(err)
		s.True(isFollow)
		wg.Done()
	}()
	wg.Wait()
}

func (s *testRelation) TestFollow() {
	err := Follow(s.ctx, 3, 4)
	s.NoError(err)

	isFollow, err := IsFollow(s.ctx, 3, 4)
	s.NoError(err)
	s.True(isFollow)

	err = Follow(s.ctx, 3, 4)
	s.NoError(err)

	res, err := ListRelation(s.ctx, 3, FOLLOW)
	s.NoError(err)
	s.EqualValues(1, len(res))
	s.EqualValues(0, res[0].FollowCount)
	s.EqualValues(1, res[0].FollowerCount)

	err = Follow(s.ctx, 4, 3)
	s.NoError(err)

	res, err = ListRelation(s.ctx, 4, FANS)
	s.NoError(err)
	s.EqualValues(1, len(res))
	s.EqualValues(1, res[0].FollowerCount)
	s.EqualValues(1, res[0].FollowCount)
}

func (s *testRelation) TestUnFollow() {
	err := UnFollow(s.ctx, 7, 8)
	s.NoError(err)

	isFollow, err := IsFollow(s.ctx, 7, 8)
	s.NoError(err)
	s.False(isFollow)

	err = UnFollow(s.ctx, 7, 8)
	s.NoError(err)
}

func (s *testRelation) TestListRelation() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		res, err := ListRelation(s.ctx, 1, FOLLOW)
		s.NoError(err)
		s.Equal(1, len(res))
		s.EqualValues(5, res[0].Id)
		wg.Done()
	}()

	go func() {
		res, err := ListRelation(s.ctx, 5, FANS)
		s.NoError(err)
		s.Equal(2, len(res))
		s.EqualValues(6^1, res[0].Id^res[1].Id)
		wg.Done()
	}()

	go func() {
		res, err := ListRelation(s.ctx, 5, FRIENDS)
		s.NoError(err)
		s.Equal(1, len(res))
		s.EqualValues(6, res[0].Id)
		wg.Done()
	}()

	wg.Wait()
}

func (s *testRelation) TestListRelationWithUserFollow() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		res, err := ListRelationWithUserFollow(s.ctx, 5, 5, FANS)
		s.NoError(err)
		s.Equal(1, len(res))
		s.EqualValues(6, res[0].Id)
		wg.Done()
	}()

	go func() {
		res, err := ListRelationWithUserFollow(s.ctx, 6, 1, FOLLOW)
		s.NoError(err)
		s.Equal(1, len(res))
		s.EqualValues(5, res[0].Id)
		wg.Done()
	}()

	go func() {
		res, err := ListRelationWithUserFollow(s.ctx, 1, 6, FRIENDS)
		s.NoError(err)
		s.Equal(1, len(res))
		s.EqualValues(5, res[0].Id)
		wg.Done()
	}()

	wg.Wait()
}

func (s *testRelation) SetupSuite() {
	initNeo4j()

	s.ctx = context.Background()
	session := Driver.NewSession(s.ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	_, err := session.Run(s.ctx, `
		MATCH (a:User)
		WHERE a.name="aya" OR a.name="satori" OR a.name="remilia" OR a.name="koishi"
		OR a.name="reimu" OR a.name="marisa" OR a.name="momiji" OR a.name="hatate"
		DETACH DELETE a`,
		map[string]any{})
	if err != nil {
		println(err.Error())
		return
	}

	_, err = session.ExecuteWrite(s.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(s.ctx, `
			CREATE (c:User{user_id:1,follower_count:0,follow_count:0,name:"aya"}),
			(:User{user_id:2,follower_count:0,follow_count:0,name:"satori"}),
			(:User{user_id:3,follower_count:0,follow_count:0,name:"remilia"}),
			(:User{user_id:4,follower_count:0,follow_count:0,name:"koishi"}),
			(a:User{user_id:5,follower_count:0,follow_count:0,name:"reimu"})
			-[:follow]->(b:User{user_id:6,follower_count:0,follow_count:0,name:"marisa"}),
			(:User{user_id:7,follower_count:0,follow_count:0,name:"momiji"})
			-[:follow]->(:User{user_id:8,follower_count:0,follow_count:0,name:"hatate"})
			MERGE (b)-[:follow]->(a)
			MERGE (c)-[:follow]->(a)`,
			map[string]any{})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		println(err.Error())
		return
	}
	s.session = session
}

func (s *testRelation) TearDownSuite() {
	_, err := s.session.Run(s.ctx, `
		MATCH (a:User)
		WHERE a.name="aya" OR a.name="satori" OR a.name="remilia" OR a.name="koishi"
		OR a.name="reimu" OR a.name="marisa" OR a.name="momiji" OR a.name="hatate"
		DETACH DELETE a`,
		map[string]any{})
	if err != nil {
		println(err.Error())
		return
	}

	err = s.session.Close(s.ctx)
	if err != nil {
		println(err.Error())
		return
	}
}

func TestRelationDb(t *testing.T) {
	suite.Run(t, new(testRelation))
}
