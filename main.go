package main

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/handler"
	"github.com/Sirok47/TOP_GAMES/repository"
	"github.com/Sirok47/TOP_GAMES/service"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	con:= &handler.TopGames{}
	con.Db, _ = sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
	defer con.Db.Close()
	con.Service = &service.TopGames{Db: con.Db}
	con.Service.Rep = &repository.TopGames{Db: con.Service.Db}
	
	e := echo.New()

	e.POST("/write/:id/:name/:rating/:platform/:date", con.CreateLine)

	e.DELETE("/delete/:id", con.DeleteLine)

	e.GET("/read/:id", con.ReadLine)

	e.Logger.Fatal(e.Start(":1323"))
}



