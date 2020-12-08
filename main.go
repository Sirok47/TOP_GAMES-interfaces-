package main

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/handler"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	con := &handler.TopGames{}
	DBConnection, _ := sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
	defer DBConnection.Close()
	handler.NewHandler(DBConnection)

	e := echo.New()

	e.POST("/write", con.CreateLine)

	e.DELETE("/delete/:id", con.DeleteLine)

	e.GET("/read/:id", con.ReadLine)

	e.PUT("/update", con.UpdateLine)

	e.Logger.Fatal(e.Start(":1323"))
}
