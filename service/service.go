package service

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/repository"
)

type TopGames struct {
	Db  *sql.DB
	Rep *repository.TopGames
}

func (S TopGames) ReadLine(id int) (*model.SingleGame,error) {
	return S.Rep.ReadLine(id)
}
func (S TopGames) CreateLine(g *model.SingleGame) error {
	return S.Rep.CreateLine(g)
}
func (S TopGames) UpdateLine(g *model.SingleGame) error {
	return S.Rep.UpdateLine(g)
}
func (S TopGames) DeleteLine(id int) error {
	return S.Rep.DeleteLine(id)
}
