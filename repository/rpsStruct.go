// Package repository works with DB
package repository

import (
	"context"

	"github.com/Sirok47/TOP_GAMES/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type CRUDforDB interface {
	Read(id int) (*model.SingleGame, error)
	Create(g *model.SingleGame) error
	Delete(id int) error
	Update(g *model.SingleGame) error
}

// TopGames stores DB connection's, context's and next structure's objects for handler package
type TopGames struct {
	db  *mongo.Collection
	ctx context.Context
}

// NewRepository is a constructor for creating "TopGames"'s object in repository package
func NewRepository(ctx context.Context, db *mongo.Collection) CRUDforDB {
	return &TopGames{db, ctx}
}
