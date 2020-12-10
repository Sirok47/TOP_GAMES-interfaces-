package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/service"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopGames struct {
	db   *mongo.Collection
	ctx  context.Context
	serv *service.TopGames
}

func NewHndl(ctx context.Context, db *mongo.Collection) *TopGames {
	return &TopGames{db, ctx, service.NewSrv(ctx, db)}
}

func (con *TopGames) Read(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	g, err := con.serv.Read(id)
	if err != nil {
		return errors.Wrap(err, "Read failed")
	}

	return c.JSON(http.StatusOK, g)
}

func (con *TopGames) Create(c echo.Context) error {
	var err error

	g := new(model.SingleGame)

	if err = c.Bind(g); err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	err = con.serv.Create(g)

	if err != nil {
		return errors.Wrap(err, "Create failed")
	}

	return c.String(http.StatusCreated, "Line have been created")
}

func (con *TopGames) Update(c echo.Context) error {
	var err error

	g := new(model.SingleGame)
	if err = c.Bind(g); err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	err = con.serv.Update(g)

	if err != nil {
		return errors.Wrap(err, "Update failed")
	}

	return c.String(http.StatusCreated, "Line have been updated")
}

func (con *TopGames) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	err = con.serv.Delete(id)

	if err != nil {
		return errors.Wrap(err, "Delete failed")
	}

	return c.String(http.StatusOK, "Line have been deleted")
}
