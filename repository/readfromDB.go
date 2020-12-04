package repository

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (g *SingleGame) Read (c echo.Context) error {
	g= &SingleGame{0,"---",0,"---","---"}
	db, _ := sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
	defer db.Close()
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := db.Query("select * from TopGames where id = $1",id)
	if err != nil {
		panic(err)
	}
	for res.Next(){
		err =res.Scan(&g.Id,&g.Name, &g.Rating, &g.Platform, &g.Date)
		if err != nil {
			panic(err)
		}
	}
	res.Close()

	return c.JSON(http.StatusOK, g)

}