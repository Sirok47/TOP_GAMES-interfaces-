package main

import (
	"context"
	"time"

	"github.com/Sirok47/TOP_GAMES_srv-rps/srv+rps/repository"
	"github.com/Sirok47/TOP_GAMES_srv-rps/srv+rps/service"

	"github.com/Sirok47/TOP_GAMES-interfaces-/handler"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbChoice = 1 // 0 for Mongo, 1 for Redis
	timeout  = 10
)

func main() {
	var (
		ctx        context.Context   = nil
		collection *mongo.Collection = nil
		conn       redis.Conn        = nil
		err        error
		con        *handler.TopGames
	)
	switch dbChoice {
	case 0:
		var cancel context.CancelFunc
		client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		ctx, cancel = context.WithTimeout(context.Background(), timeout*time.Second)

		defer cancel()

		_ = client.Connect(ctx)
		collection = client.Database("TOP_GAMES").Collection("TopGames")
		ctx = context.TODO()
		con = handler.NewHandler(service.NewService(repository.NewMongoRepository(ctx, collection)))
	case 1:
		conn, err = redis.Dial("tcp", "localhost:6379")
		if err != nil {
			return
		}
		defer func() {
			err := conn.Close()
			if err != nil {
				return
			}
		}()
		con = handler.NewHandler(service.NewService(repository.NewRedisRepository(conn)))
	default:
		return
	}

	e := echo.New()

	e.POST("/write", con.Create)

	e.DELETE("/delete/:id", con.Delete)

	e.GET("/read/:id", con.Read)

	e.PUT("/update", con.Update)

	e.Logger.Fatal(e.Start(":1323"))
}
