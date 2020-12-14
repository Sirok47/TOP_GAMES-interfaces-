// Package repository works with DB
package repository

import (
	"context"

	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

// DBTemplate for mongo and redis
type DBTemplate interface {
	Create(g *model.SingleGame) error
	Read(id int) (*model.SingleGame, error)
	Update(g *model.SingleGame) error
	Delete(id int) error
}

// TopGamesMongo stores DB connection's and context's objects for mongoDB
type TopGamesMongo struct {
	db  *mongo.Collection
	ctx context.Context
}

// TopGamesRedis stores DB connection's object for redis
type TopGamesRedis struct {
	db redis.Conn
}

// NewRepository is a constructor for creating "TopGames"'s object in repository package
func NewRepository(ctx context.Context, dbMongo *mongo.Collection, dbRedis redis.Conn) DBTemplate {
	if ctx != nil {
		return &TopGamesMongo{dbMongo, ctx}
	}
	return &TopGamesRedis{dbRedis}
}
