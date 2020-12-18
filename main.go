package main

import (
	grpcpb "github.com/Sirok47/TOP_GAMES-interfaces-/grpc"
	"github.com/Sirok47/TOP_GAMES-interfaces-/handler"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func main() {
	gcon, _ := grpc.Dial(":8080", grpc.WithInsecure())
	cli := grpcpb.NewCRUDClient(gcon)
	con := handler.NewHandler(cli)

	e := echo.New()

	e.POST("/write", con.Create)

	e.DELETE("/delete/:id", con.Delete)

	e.GET("/read/:id", con.Read)

	e.PUT("/update", con.Update)

	e.Logger.Fatal(e.Start(":1323"))
}
