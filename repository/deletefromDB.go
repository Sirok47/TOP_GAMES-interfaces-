package repository

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Delete(c echo.Context,db *sql.DB) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_,err := db.Exec("delete from TopGames where id = $1", id)
	if err != nil{
		panic(err)
	}
	return
}