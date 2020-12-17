// Package handler contains function for collecting data from requests
package handler

import (
	"context"
	"net/http"
	"strconv"

	grpcpb "github.com/Sirok47/TOP_GAMES-interfaces-/grpc"

	"github.com/Sirok47/TOP_GAMES_srv-rps/srv+rps/service"

	"github.com/Sirok47/TOP_GAMES-interfaces-/model"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// TopGames stores DB connection's, context's and next structure's objects for handler package
type TopGames struct {
	srv *service.TopGames
	cli grpcpb.CRUDClient
}

// NewHandler is a constructor for creating "TopGames"'s object in handler package
func NewHandler(srv *service.TopGames, cli grpcpb.CRUDClient) *TopGames {
	return &TopGames{srv, cli}
}

// Read gets id from request and passes in to srv.Read
func (con *TopGames) Read(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	g, err := con.cli.Read(context.Background(), &grpcpb.Id{ID: int32(id)})
	if err != nil {
		return errors.Wrap(err, "Read failed")
	}

	return c.JSON(http.StatusOK, g)
}

// Create decodes JSON request into "TopGames"'s object and passes it to srv.Create
func (con *TopGames) Create(c echo.Context) error {
	g := new(model.SingleGame)

	if err := c.Bind(g); err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	err, _ := con.cli.Create(context.Background(), &grpcpb.Structmsg{ID: int32(g.ID), Name: g.Name, Rating: int32(g.Rating), Platform: g.Platform, Date: g.Date})
	if err.Err != "" {
		return c.String(http.StatusInternalServerError, err.Err)
	}
	return c.String(http.StatusCreated, "Line have been created")
}

// Update decodes JSON request into "TopGames"'s object and passes it to srv.Update
func (con *TopGames) Update(c echo.Context) error {
	g := new(model.SingleGame)
	if err := c.Bind(g); err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	err, _ := con.cli.Update(context.Background(), &grpcpb.Structmsg{ID: int32(g.ID), Name: g.Name, Rating: int32(g.Rating), Platform: g.Platform, Date: g.Date})

	if err.Err != "" {
		return c.String(http.StatusInternalServerError, err.Err)
	}
	return c.String(http.StatusCreated, "Line have been updated")
}

// Delete gets id from request and passes in to srv.Delete
func (con *TopGames) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.Wrap(err, "Wrong input")
	}

	err2, _ := con.cli.Delete(context.Background(), &grpcpb.Id{ID: int32(id)})

	if err2.Err != "" {
		return c.String(http.StatusInternalServerError, err2.Err)
	}
	return c.String(http.StatusOK, "Line have been deleted")
}
