package main

import (
	"context"
	"time"

	"github.com/Sirok47/TOP_GAMES/handler"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

const timeout = 10

func main() {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)

	defer cancel()

	_ = client.Connect(ctx)
	collection := client.Database("TOP_GAMES").Collection("TopGames")
	ctx = context.TODO()
	con := handler.NewHndl(ctx, collection)

	e := echo.New()

	e.POST("/write", con.Create)

	e.DELETE("/delete/:id", con.Delete)

	e.GET("/read/:id", con.Read)

	e.PUT("/update", con.Update)

	e.Logger.Fatal(e.Start(":1323"))
}
