package main

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/handler"
	"github.com/Sirok47/TOP_GAMES/repository"
	"github.com/Sirok47/TOP_GAMES/service"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)
type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}
func main() {
	con := &handler.TopGames{}
	con.Db, _ = sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
	defer con.Db.Close()
	con.Service = &service.TopGames{Db: con.Db}
	con.Service.Rep = &repository.TopGames{Db: con.Service.Db}

	e := echo.New()

	e.POST("/write", con.CreateLine)

	e.DELETE("/delete/:id", con.DeleteLine)

	e.GET("/read/:id", con.ReadLine)

	e.PUT("/update", con.UpdateLine)

	e.Logger.Fatal(e.Start(":1323"))
}
