package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (R TopGames) UpdateLine(g *model.SingleGame) error {
	_, err := R.db.ReplaceOne(R.ctx,bson.D{primitive.E{Key: "_id", Value: g.Id}},bson.M{"_id":g.Id,"Name": g.Name,"Rating":g.Rating,"Platform":g.Platform,"Date":g.Date})
	return err
}
