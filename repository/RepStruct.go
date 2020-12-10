package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopGames struct {
	db *mongo.Collection
	ctx context.Context
}
func NewRep(db *mongo.Collection, ctx context.Context) *TopGames {
	return &TopGames{db,ctx}
}