package service

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/repository"
)
type TopGames struct {
	Db *sql.DB
}
func (S TopGames) ReadLine(id int) *model.SingleGame {
	Rep := repository.TopGames{Db:S.Db}
	return Rep.ReadLine(id)
}
func (S TopGames) CreateLine(g *model.SingleGame) {
	Rep := repository.TopGames{Db:S.Db}
	Rep.CreateLine(g)
}
func (S TopGames) DeleteLine(id int) {
	Rep := repository.TopGames{Db:S.Db}
	Rep.DeleteLine(id)
}
