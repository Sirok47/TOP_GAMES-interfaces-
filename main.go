package main

import (
	"TOP_GAMES/repository"
	"TOP_GAMES/service"
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)


func main() {
	g:=repository.SingleGame{}

	e := echo.New()

	db, _ := sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
	defer db.Close()

	e.POST("/write/:id/:name/:rating/:platform/:date", g.Write(e,db))

	e.DELETE("/delete/:id", repository.Delete(e,db))

	e.GET("/read/:id", g.Read(c,db))

	e.Logger.Fatal(e.Start(":1323"))
}



