package handler

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"TOP_GAMES/service"
)

func Readhandler(c echo.Context, db *sql.DB) {
	return service.Readservice(c,db)
}
func Writehandler(c echo.Context, db *sql.DB) {
	return service.Writeservice(c,db)
}
func Deletehandler(c echo.Context, db *sql.DB) {
	return service.Deleteservice(c,db)
}
