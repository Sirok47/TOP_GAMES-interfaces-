package handler

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/service"
	"github.com/labstack/echo/v4"
	"strconv"
)
type TopGamesHandler struct {
	Db *sql.DB
}

func (con *TopGamesHandler) Readhandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	g := model.SingleGame{}
	Service := service.TopGamesService{Db:con.Db}
	Service.Readservice(&g,id)
	return c.JSON(200, g)
}
func (con *TopGamesHandler) Writehandler(c echo.Context) error {
	g := model.SingleGame{}
	g.Id,_ = strconv.Atoi(c.Param("id"))
	g.Name = c.Param("name")
	rr, _ := strconv.Atoi(c.Param("rating"))
	g.Rating=float64(rr)
	g.Platform= c.Param("platform")
	g.Date= c.Param("date")
	Service := service.TopGamesService{Db:con.Db}
	Service.Writeservice(&g)
	return c.String(200, "Done!")
}
func (con *TopGamesHandler) Deletehandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Service := service.TopGamesService{Db:con.Db}
	Service.Deleteservice(id)
	return c.String(200, "Line have been deleted.")
}

