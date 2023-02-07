package db

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestIsFollow(t *testing.T) {
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		isFollow, err := IsFollow(ctx, 1, 2)
		assert.Nil(t, err)
		assert.False(t, isFollow)
		wg.Done()
	}()
	go func() {
		isFollow, err := IsFollow(ctx, 5, 6)
		assert.Nil(t, err)
		assert.True(t, isFollow)
		wg.Done()
	}()
	wg.Wait()
}

func TestFollow(t *testing.T) {
	ctx := context.Background()
	err := Follow(ctx, 3, 4)
	assert.Nil(t, err)

	isFollow, err := IsFollow(ctx, 3, 4)
	assert.Nil(t, err)
	assert.True(t, isFollow)

	err = Follow(ctx, 3, 4)
	assert.Nil(t, err)
}

func TestUnFollow(t *testing.T) {
	ctx := context.Background()
	err := UnFollow(ctx, 7, 8)
	assert.Nil(t, err)

	isFollow, err := IsFollow(ctx, 7, 8)
	assert.Nil(t, err)
	assert.False(t, isFollow)

	err = UnFollow(ctx, 7, 8)
	assert.Nil(t, err)
}

func TestListRelation(t *testing.T) {
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		res, err := ListRelation(ctx, 1, FOLLOW)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, uint32(5), res[0].Id)
		wg.Done()
	}()

	go func() {
		res, err := ListRelation(ctx, 5, FANS)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(res))
		assert.Equal(t, uint32(6^1), res[0].Id^res[1].Id)
		wg.Done()
	}()

	go func() {
		res, err := ListRelation(ctx, 5, FRIENDS)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, uint32(6), res[0].Id)
		wg.Done()
	}()

	wg.Wait()
}

func TestListRelationWithUserFollow(t *testing.T) {
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		res, err := ListRelationWithUserFollow(ctx, 5, 5, FANS)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, uint32(6), res[0].Id)
		wg.Done()
	}()

	go func() {
		res, err := ListRelationWithUserFollow(ctx, 6, 1, FOLLOW)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, uint32(5), res[0].Id)
		wg.Done()
	}()

	go func() {
		res, err := ListRelationWithUserFollow(ctx, 1, 6, FRIENDS)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, uint32(5), res[0].Id)
		wg.Done()
	}()

	wg.Wait()
}

func TestMain(m *testing.M) {
	Init()

	ctx := context.Background()
	session := Driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	_, err := session.Run(ctx, `
		MATCH (a:User)
		WHERE a.name="aya" OR a.name="satori" OR a.name="remilia" OR a.name="koishi"
		OR a.name="reimu" OR a.name="marisa" OR a.name="momiji" OR a.name="hatate"
		DETACH DELETE a`,
		map[string]any{})
	if err != nil {
		println(err.Error())
		return
	}

	_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
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

	m.Run()

	_, err = session.Run(ctx, `
		MATCH (a:User)
		WHERE a.name="aya" OR a.name="satori" OR a.name="remilia" OR a.name="koishi"
		OR a.name="reimu" OR a.name="marisa" OR a.name="momiji" OR a.name="hatate"
		DETACH DELETE a`,
		map[string]any{})
	if err != nil {
		println(err.Error())
		return
	}

	err = session.Close(ctx)
	if err != nil {
		println(err.Error())
		return
	}
}
