package service

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"TOP_GAMES/repository"
)

func Readservice(c echo.Context, db *sql.DB) {
	return g.Read(c,db)
}
func Writehandler(c echo.Context, db *sql.DB) {
	return g.Write(c,db)
}
func Deletehandler(c echo.Context, db *sql.DB) {
	return repository.Delete(c,db)
}
