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
	Service := service.TopGamesService{Db:con.Db}
	Service.Readservice(&model.G,id)
	return c.JSON(200, model.G)
}
func (con *TopGamesHandler) Writehandler(c echo.Context) error {
	model.G.Id,_ = strconv.Atoi(c.Param("id"))
	model.G.Name = c.Param("name")
	rr, _ := strconv.Atoi(c.Param("rating"))
	model.G.Rating=float64(rr)
	model.G.Platform= c.Param("platform")
	model.G.Date= c.Param("date")
	Service := service.TopGamesService{Db:con.Db}
	Service.Writeservice()
	return c.String(200, "Done!")
}
func (con *TopGamesHandler) Deletehandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Service := service.TopGamesService{Db:con.Db}
	Service.Deleteservice(id)
	return c.String(200, "Line have been deleted.")
}

