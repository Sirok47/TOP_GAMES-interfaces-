package repository

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Write (db *sql.DB, g *SingleGame) {
	_,err:=db.Exec("insert into TopGames (id,GameName,Rating,Platform,ReleaseDate) values ($1,$2,$3,$4,$5)",g.Id,g.Name,g.Rating,g.Platform,g.Date)
	if err != nil{
		panic(err)
	}
	return
}