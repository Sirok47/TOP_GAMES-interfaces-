package main

import (
	"context"
	"github.com/Sirok47/TOP_GAMES/handler"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Connect(ctx)
	collection := client.Database("TOP_GAMES").Collection("TopGames")
	ctx =context.TODO()
	con:=handler.NewHandler(collection,ctx)

	e := echo.New()

	e.POST("/write", con.CreateLine)

	e.DELETE("/delete/:id", con.DeleteLine)

	e.GET("/read/:id", con.ReadLine)

	e.PUT("/update", con.UpdateLine)

	e.Logger.Fatal(e.Start(":1323"))
}
