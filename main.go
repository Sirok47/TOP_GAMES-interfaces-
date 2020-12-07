package main

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/handler"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	con:= &handler.TopGamesHandler{}

	con.Db, _ = sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
	defer con.Db.Close()
	
	e := echo.New()

	e.POST("/write/:id/:name/:rating/:platform/:date", con.Writehandler)

	e.DELETE("/delete/:id", con.Deletehandler)

	e.GET("/read/:id", con.Readhandler)

	e.Logger.Fatal(e.Start(":1323"))
}



