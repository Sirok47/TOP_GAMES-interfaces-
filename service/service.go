package service

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"TOP_GAMES/repository"
	. "TOP_GAMES/model"
)

func (con TopGamesHandler) Readservice() {
	return Read(db,g)
}
func Writeservice(con TopGamesHandler) {
	return Write(db,g)
}
func Deleteservice(db *sql.DB, g SingleGame) {
	return Delete(db,g)
}
