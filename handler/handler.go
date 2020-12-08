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
	id, err := strconv.Atoi(c.Param("id"))
	if err!=nil{return c.String(500, "ID can't be converted")}
	return c.JSON(200, con.Service.ReadLine(id))
}
func (con *TopGames) CreateLine(c echo.Context) error {
	g := model.SingleGame{}
	rr,err := strconv.Atoi(c.Param("rating"))
	if err!=nil{return c.String(500, "Rating can't be converted")}
	g.Rating=float64(rr)
	g.Id,err = strconv.Atoi(c.Param("id"))
	if err!=nil{return c.String(500, "ID can't be converted")}
	g.Name = c.Param("name")
	g.Platform= c.Param("platform")
	g.Date= c.Param("date")
	return c.String(201, con.Service.CreateLine(&g))
}
func (con *TopGames) DeleteLine(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err!=nil{return c.String(500, "ID can't be converted")}
	return c.String(200, con.Service.DeleteLine(id))
}

