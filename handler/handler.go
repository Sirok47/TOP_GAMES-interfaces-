package handler

import (
	"context"
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type TopGames struct {
	db   *mongo.Collection
	ctx  context.Context
	serv *service.TopGames
}

func NewHandler(db *mongo.Collection, ctx context.Context) *TopGames {
	return &TopGames{db, ctx, service.NewService(db, ctx)}
}

func (con *TopGames) ReadLine(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	g, err := con.serv.ReadLine(id)
	if err != nil {
		return err
	}
	return c.JSON(200, g)
}
func (con *TopGames) CreateLine(c echo.Context) error {
	var err error
	g := new(model.SingleGame)
	if err = c.Bind(g); err != nil {
		return err
	}
	err = con.serv.CreateLine(g)
	if err != nil {
		return err
	}
	return c.String(201, "Line have been created")
}
func (con *TopGames) UpdateLine(c echo.Context) error {
	var err error
	g := new(model.SingleGame)
	if err = c.Bind(g); err != nil {
		return err
	}
	err = con.serv.UpdateLine(g)
	if err != nil {
		return err
	}
	return c.String(201, "Line have been updated")
}
func (con *TopGames) DeleteLine(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = con.serv.DeleteLine(id)
	if err != nil {
		return err
	}
	return c.String(200, "Line have been deleted")
}
