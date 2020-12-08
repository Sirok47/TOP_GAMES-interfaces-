package handler

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/service"
	"github.com/labstack/echo/v4"
	"strconv"
)
type TopGames struct {
	Db *sql.DB
	Service *service.TopGames
}

func (con *TopGames) ReadLine(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(200, con.Service.ReadLine(id))
}
func (con *TopGames) CreateLine(c echo.Context) error {
	g := model.SingleGame{}
	g.Id,_ = strconv.Atoi(c.Param("id"))
	g.Name = c.Param("name")
	rr, _ := strconv.Atoi(c.Param("rating"))
	g.Rating=float64(rr)
	g.Platform= c.Param("platform")
	g.Date= c.Param("date")
	return c.String(200, con.Service.CreateLine(&g))
}
func (con *TopGames) DeleteLine(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.String(200, con.Service.DeleteLine(id))
}

