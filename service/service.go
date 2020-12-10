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
func NewService(db *mongo.Collection,ctx context.Context) *TopGames {
	return &TopGames{db,ctx, repository.NewRep(db,ctx)}
}

func (S TopGames) ReadLine(id int) (*model.SingleGame, error) {
	return S.rep.ReadLine(id)
}
func (S TopGames) CreateLine(g *model.SingleGame) error {
	return S.rep.CreateLine(g)
}
func (S TopGames) UpdateLine(g *model.SingleGame) error {
	return S.rep.UpdateLine(g)
}
func (S TopGames) DeleteLine(id int) error {
	return S.rep.DeleteLine(id)
}
