package service

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/repository"
)
type TopGamesHandler struct {
	Db *sql.DB
}
func Readservice(db *sql.DB,g *model.SingleGame,id int) {
	repository.Read(db,g,id)
}
func Writeservice(db *sql.DB) {
	repository.Write(db)
}
func Deleteservice(db *sql.DB,id int) {
	repository.Delete(db,id)
}
