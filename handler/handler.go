package handler

import (
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/service"
	"github.com/labstack/echo/v4"
	"strconv"
)

func (con *model.TopGamesHandler) Readhandler(c echo.Context) error {
	con.Readservice()
	return c.JSON(200, con.G)
}
func (con *model.TopGamesHandler) Writehandler(c echo.Context) error {
	con.G.Id,_ = strconv.Atoi(c.Param("id"))
	con.G.Name = c.Param("name")
	rr, _ := strconv.Atoi(c.Param("rating"))
	con.G.Rating=float64(rr)
	con.G.Platform= c.Param("platform")
	con.G.Date= c.Param("date")
	Writeservice(con)
	return c.String(200, "Done!")
}
func (con *model.TopGamesHandler) Deletehandler(c echo.Context) error {
	Deleteservice(con.Db,con.G)
	return c.String(200, "Line have been deleted.")
}

