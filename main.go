package main

import (
	"TOP_GAMES/repository"
	"TOP_GAMES/service"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)


func main() {
	g:=repository.SingleGame{}

	e := echo.New()

	service.PostDelete(g,e)

	e.GET("/read/:id", g.Read)

	e.Logger.Fatal(e.Start(":1323"))
}



