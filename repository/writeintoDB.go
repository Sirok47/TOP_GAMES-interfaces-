package repository

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (g *SingleGame) Write(c echo.Context,db *sql.DB) error {
	g.Id,_ = strconv.Atoi(c.Param("id"))
	g.Name = c.Param("name")
	rr, _ := strconv.Atoi(c.Param("rating"))
	g.Rating=float64(rr)
	g.Platform= c.Param("platform")
	g.Date= c.Param("date")
	_,err:=db.Exec("insert into TopGames (id,GameName,Rating,Platform,ReleaseDate) values ($1,$2,$3,$4,$5)",g.Id,g.Name,g.Rating,g.Platform,g.Date)
	if err != nil{
		panic(err)
	}
	return c.String(http.StatusOK, "Done!")
}