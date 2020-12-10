package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (R TopGames) ReadLine(id int) (*model.SingleGame, error) {
	g := &model.SingleGame{Id: id, Name: "---", Rating: 0, Platform: "---", Date: "---"}
	err := R.db.FindOne(R.ctx, bson.D{primitive.E{Key: "_id", Value: g.Id}}).Decode(&g)
	return g, err
}
