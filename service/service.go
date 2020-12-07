package service

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/repository"
)
type TopGamesService struct {
	Db *sql.DB
}
func (S TopGamesService) Readservice(g *model.SingleGame,id int) {
	Rep := repository.TopGamesRepository{Db:S.Db}
	Rep.Read(g,id)
}
func (S TopGamesService) Writeservice(g *model.SingleGame) {
	Rep := repository.TopGamesRepository{Db:S.Db}
	Rep.Write(g)
}
func (S TopGamesService) Deleteservice(id int) {
	Rep := repository.TopGamesRepository{Db:S.Db}
	Rep.Delete(id)
}
