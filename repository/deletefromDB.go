package repository

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Delete(c echo.Context) error {
	db, _ := sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
	defer db.Close()
	id, _ := strconv.Atoi(c.Param("id"))
	_,err := db.Exec("delete from TopGames where id = $1", id)
	if err != nil{
		panic(err)
	}
	return c.String(http.StatusOK, "Line have been deleted.")
}