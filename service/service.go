package service

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/repository"
)

type TopGames struct {
	db  *sql.DB
	rep *repository.TopGames
}
func NewService(db *sql.DB) *TopGames {
	return &TopGames{db,repository.NewRep(db)}
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
