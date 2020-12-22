// Package handler contains function for collecting data from requests
package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	grpcpb "github.com/Sirok47/TOP_GAMES-interfaces-/grpc"

	"github.com/Sirok47/TOP_GAMES-interfaces-/model"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// TopGames stores DB connection's, context's and next structure's objects for handler package
type TopGames struct {
	cli   grpcpb.CRUDClient
	users map[string]string
}

// NewHandler is a constructor for creating "TopGames"'s object in handler package
func NewHandler(cli grpcpb.CRUDClient, u map[string]string) *TopGames {
	return &TopGames{cli, u}
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

	con.cli.Create(context.Background(), &grpcpb.Structmsg{ID: int32(g.ID), Name: g.Name, Rating: int32(g.Rating), Platform: g.Platform, Date: g.Date})

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

func (con *TopGames) Login(c echo.Context) error {
	username := c.Param("name")
	password := c.Param("pass")
	if con.users[username] == password {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Max"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		t, err := token.SignedString([]byte("sirok"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}
