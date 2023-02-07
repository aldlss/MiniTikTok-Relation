package db

import (
	"context"
	"fmt"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/relation/model"
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/relation/pack"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/errno"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	log "github.com/sirupsen/logrus"
)

const DatabaseName = "neo4j"

type relationType int

const (
	FOLLOW relationType = iota
	FANS
	FRIENDS
)

func Follow(ctx context.Context, id uint32, toId uint32) error {
	session := Driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: DatabaseName})

	res, err := session.Run(ctx, `
		MATCH (a:User)-[c:follow]->(b:User)
		WHERE a.user_id = $id AND b.user_id = $to_id
		RETURN c`,
		map[string]any{
			"id":    id,
			"to_id": toId,
		})
	if err != nil {
		return err
	}
	// 传回是否有返回 c，若有说明已经关注过
	if res.Next(ctx) == true {
		log.Warn("关注失败，已关注过")
		return nil
	}

	tra, err := session.BeginTransaction(ctx, func(config *neo4j.TransactionConfig) {})

	res, err = tra.Run(ctx, `
		MATCH (a:User), (b:User)
		WHERE a.user_id = $id AND b.user_id = $to_id
		MERGE (a)-[r:follow]->(b)
		RETURN r`,
		map[string]any{
			"id":    id,
			"to_id": toId,
		})
	if err != nil {
		return err
	}
	// 传回是否有返回 r，若没有说明创建失败
	if res.Next(ctx) == false {
		log.Warn("not create follow relation, maybe because not find User")
		return nil
	}

	res, err = tra.Run(ctx, `
		MATCH (a:User), (b:User)
		WHERE a.user_id = $id AND b.user_id = $to_id
		SET a.follow = a.follow + 1, b.fans = b.fans + 1`,
		map[string]any{
			"id":    id,
			"to_id": toId,
		})
	if err != nil {
		err1 := tra.Rollback(ctx)
		if err1 != nil {
			log.Fatal(fmt.Sprintf("roll back fail:%v", err1.Error()))
		}
		return err
	}

	err = tra.Commit(ctx)
	if err != nil {
		err1 := tra.Rollback(ctx)
		if err1 != nil {
			log.Fatal(fmt.Sprintf("roll back fail:%v", err1.Error()))
		}
		return err
	}

	err = tra.Close(ctx)
	if err != nil {
		return err
	}

	err = session.Close(ctx)
	if err != nil {
		return err
	}

	return nil
}

func UnFollow(ctx context.Context, id uint32, toId uint32) error {
	session := Driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: DatabaseName})

	res, err := session.Run(ctx, `
		MATCH (a:User)-[c:follow]->(b:User)
		WHERE a.user_id = $id AND b.user_id = $to_id
		RETURN c`,
		map[string]any{
			"id":    id,
			"to_id": toId,
		})
	if err != nil {
		return err
	}
	// 检查是否返回 c，没有说明不存在该关系
	if res.Next(ctx) == false {
		log.Warn("取关失败，并未关注")
		return nil
	}

	res, err = session.Run(ctx, `
		MATCH (a:User)-[c:follow]->(b:User)
		WHERE a.user_id = $id AND b.user_id = $to_id
		DELETE c
		SET a.follow = a.follow - 1, b.fans = b.fans - 1`,
		map[string]any{
			"id":    id,
			"to_id": toId,
		})
	if err != nil {
		return err
	}

	err = session.Close(ctx)
	if err != nil {
		return err
	}

	return nil
}

func ListRelation(ctx context.Context, toId uint32, relationType relationType) ([]*model.User, error) {
	session := Driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: DatabaseName,
	})

	var cypher string
	switch relationType {
	case FOLLOW:
		cypher = `
			MATCH (a:User)-[:follow]->(b:User)
			WHERE a.user_id = $to_id
			RETURN b`
	case FANS:
		cypher = `
			MATCH (a:User)-[:follow]->(b:User)
			WHERE b.user_id = $to_id
			RETURN a`
	case FRIENDS:
		cypher = `
			MATCH (a:User)-[:follow]->(b:User), (b:User)-[:follow]->(a:User)
			WHERE b.user_id = $to_id
			RETURN a`
	default:
		return nil, errno.ParamErr
	}

	records, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, cypher,
			map[string]any{
				"to_id": toId,
			})
		if err != nil {
			return nil, err
		}
		// 这里直接返回 res 再在外面对其处理会无结果，似乎必须要在这里把 res 的结果用了，也许是惰性求值？
		return res.Collect(ctx)
	})
	if err != nil {
		return nil, err
	}

	return pack.DbUsers(records.([]*neo4j.Record)), nil
}

func ListRelationWithUserFollow(ctx context.Context, id uint32, toId uint32, relationType relationType) ([]*model.User, error) {
	session := Driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: DatabaseName,
	})

	var cypher string
	switch relationType {
	case FOLLOW:
		cypher = `
			MATCH (a:User)-[:follow]->(b:User), (c:User)-[:follow]->(b:User)
			WHERE a.user_id = $to_id AND c.user_id = $id
			RETURN b`
	case FANS:
		cypher = `
			MATCH (a:User)-[:follow]->(b:User), (c:User)-[:follow]->(a:User)
			WHERE b.user_id = $to_id AND c.user_id = $id
			RETURN a`
	case FRIENDS:
		cypher = `
			MATCH (a:User)-[:follow]->(b:User), (b:User)-[:follow]->(a:User), (c:User)-[:follow]->(a:User)
			WHERE b.user_id = $to_id AND c.user_id = $id
			RETURN a`
	default:
		return nil, errno.ParamErr
	}

	records, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, cypher,
			map[string]any{
				"id":    id,
				"to_id": toId,
			})
		if err != nil {
			return nil, err
		}
		return res.Collect(ctx)
	})
	if err != nil {
		return nil, err
	}

	return pack.DbUsers(records.([]*neo4j.Record)), nil
}

func IsFollow(ctx context.Context, id uint32, toId uint32) (bool, error) {
	session := Driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: DatabaseName,
	})

	isFollow, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (a:User)-[c:follow]->(b:User)
			WHERE a.user_id = $id AND b.user_id = $to_id
			RETURN c`,
			map[string]any{
				"id":    id,
				"to_id": toId,
			})
		if err != nil {
			return nil, err
		}
		return res.Next(ctx), err
	})
	if err != nil {
		return false, err
	}

	return isFollow.(bool), nil
}
