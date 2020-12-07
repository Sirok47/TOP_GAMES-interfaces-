package service

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/Sirok47/TOP_GAMES/repository"
	"github.com/Sirok47/TOP_GAMES/model"
)

func (con *model.TopGamesHandler) Readservice() {
	return Read(db,g)
}
func (con *model.TopGamesHandler) Writeservice() {
	return Write(db,g)
}
func (con *model.TopGamesHandler) Deleteservice() {
	return Delete(db,g)
}
