// Package handler contains function for collecting data from requests
package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES/repository"
	"github.com/Sirok47/TOP_GAMES/service"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

// TopGames stores DB connection's, context's and next structure's objects for handler package
type TopGames struct {
	srv *service.TopGames
}

// NewHandler is a constructor for creating "TopGames"'s object in handler package
func NewHandler(ctx context.Context, dbMongo *mongo.Collection, dbRedis redis.Conn) *TopGames {
	return &TopGames{service.NewService(repository.NewRepository(ctx, dbMongo, dbRedis))}
}

// Read gets id from request and passes in to srv.Read
func (con *TopGames) Read(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	g, err := con.srv.Read(id)
	if err != nil {
		return errors.Wrap(err, "Read failed")
	}

	return c.JSON(http.StatusOK, g)
}

// Create decodes JSON request into "TopGames"'s object and passes it to srv.Create
func (con *TopGames) Create(c echo.Context) error {
	var err error

	g := new(model.SingleGame)

	if err = c.Bind(g); err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	err = con.srv.Create(g)

	if err != nil {
		return errors.Wrap(err, "Create failed")
	}

	return c.String(http.StatusCreated, "Line have been created")
}

// Update decodes JSON request into "TopGames"'s object and passes it to srv.Update
func (con *TopGames) Update(c echo.Context) error {
	var err error

	g := new(model.SingleGame)
	if err = c.Bind(g); err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	err = con.srv.Update(g)

	if err != nil {
		return errors.Wrap(err, "Update failed")
	}

	return c.String(http.StatusCreated, "Line have been updated")
}

// Delete gets id from request and passes in to srv.Delete
func (con *TopGames) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	err = con.srv.Delete(id)

	if err != nil {
		return errors.Wrap(err, "Delete failed")
	}

	return c.String(http.StatusOK, "Line have been deleted")
}
