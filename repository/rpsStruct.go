package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type TopGames struct {
	db  *mongo.Collection
	ctx context.Context
}

func NewRps(ctx context.Context, db *mongo.Collection) *TopGames {
	return &TopGames{db, ctx}
}
