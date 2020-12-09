package handler

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/service"
	"github.com/labstack/echo/v4"
	"strconv"
)

type TopGames struct {
	db   *sql.DB
	serv *service.TopGames
}

func NewHandler(db *sql.DB) *TopGames {
	return &TopGames{db, service.NewService(db)}
}

func (con *TopGames) ReadLine(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, err)
	}
	g, err := con.serv.ReadLine(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, g)
}
func (con *TopGames) CreateLine(c echo.Context) error {
	var err error
	g := new(model.SingleGame)
	gg := new(model.JSON)
	if err = c.Bind(gg); err != nil {
		return c.JSON(500, err)
	}
	g.Id, err = strconv.Atoi(gg.Id)
	if err != nil {
		return c.JSON(500, err)
	}
	g.Rating, err = strconv.ParseFloat(gg.Rating, 64)
	if err != nil {
		return c.JSON(500, err)
	}
	g.Name = gg.Name
	g.Date = gg.Date
	g.Platform = gg.Platform
	err = con.serv.CreateLine(g)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.String(201, "Line have been created")
}
func (con *TopGames) UpdateLine(c echo.Context) error {
	var err error
	g := new(model.SingleGame)
	gg := new(model.JSON)
	if err = c.Bind(gg); err != nil {
		return c.JSON(500, err)
	}
	g.Id, err = strconv.Atoi(gg.Id)
	if err != nil {
		return c.JSON(500, err)
	}
	g.Rating, err = strconv.ParseFloat(gg.Rating, 64)
	if err != nil {
		return c.JSON(500, err)
	}
	g.Name = gg.Name
	g.Date = gg.Date
	g.Platform = gg.Platform
	err = con.serv.UpdateLine(g)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.String(201, "Line have been updated")
}
func (con *TopGames) DeleteLine(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, err)
	}
	err = con.serv.DeleteLine(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.String(200, "Line have been deleted")
}
