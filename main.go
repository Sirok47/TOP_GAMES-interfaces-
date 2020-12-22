package main

import (
	grpcpb "github.com/Sirok47/TOP_GAMES-interfaces-/grpc"
	"github.com/Sirok47/TOP_GAMES-interfaces-/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func main() {
	users := map[string]string{"max": "5742368M"}
	LoggedIn := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("sirok"),
	})
	gcon, _ := grpc.Dial(":8080", grpc.WithInsecure())
	cli := grpcpb.NewCRUDClient(gcon)
	con := handler.NewHandler(cli, users)

	e := echo.New()

	e.POST("/login/:name/:pass", con.Login)

	e.POST("/write", con.Create, LoggedIn)

	e.DELETE("/delete/:id", con.Delete, LoggedIn)

	e.GET("/read/:id", con.Read, LoggedIn)

	e.PUT("/update", con.Update, LoggedIn)

	e.Logger.Fatal(e.Start(":1323"))
}
