package service

import (
	"context"

	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopGames struct {
	db  *mongo.Collection
	ctx context.Context
	rep *repository.TopGames
}

func NewSrv(ctx context.Context, db *mongo.Collection) *TopGames {
	return &TopGames{db, ctx, repository.NewRps(ctx, db)}
}

func (s TopGames) Read(id int) (*model.SingleGame, error) {
	return s.rep.Read(id)
}

func (s TopGames) Create(g *model.SingleGame) error {
	return s.rep.Create(g)
}

func (s TopGames) Update(g *model.SingleGame) error {
	return s.rep.Update(g)
}

func (s TopGames) Delete(id int) error {
	return s.rep.Delete(id)
}
