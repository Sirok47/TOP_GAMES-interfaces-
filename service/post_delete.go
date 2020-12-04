package service

import (
	"TOP_GAMES/repository"
	"github.com/labstack/echo/v4"
)

func PostDelete(g repository.SingleGame,e *echo.Echo) {

	e.POST("/write/:id/:name/:rating/:platform/:date", g.Write)

	e.DELETE("/delete/:id", repository.Delete)
}
