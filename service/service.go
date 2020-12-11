// Package service just passes values to repository
package service

import (
	"context"

	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/repository"
	//"github.com/Sirok47/TOP_GAMES/repository/mongoDB"
	"go.mongodb.org/mongo-driver/mongo"
)

// TopGames stores DB connection's, context's and next structure's objects for service package
type TopGames struct {
	db  *mongo.Collection
	ctx context.Context
	rps repository.CRUDforDB
}

// NewService is a constructor for creating "TopGames"'s object in service package
func NewService(ctx context.Context, db *mongo.Collection) *TopGames {
	return &TopGames{db, ctx, repository.NewRepository(ctx, db)}
}

// Read passes id to rps.Read
func (s TopGames) Read(id int) (*model.SingleGame, error) {
	return s.rps.Read(id)
}

// Create passes "TopGames"'s object to rps.Create
func (s TopGames) Create(g *model.SingleGame) error {
	return s.rps.Create(g)
}

// Update passes "TopGames"'s object to rps.Update
func (s TopGames) Update(g *model.SingleGame) error {
	return s.rps.Update(g)
}

// Delete passes id to rps.Delete
func (s TopGames) Delete(id int) error {
	return s.rps.Delete(id)
}
