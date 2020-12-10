package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (R TopGames) CreateLine(g *model.SingleGame) error {
	_, err := R.db.InsertOne(R.ctx,bson.M{"_id":g.Id,"Name":g.Name,"Rating":g.Rating,"Platform":g.Platform,"Date":g.Date})
	return err
}
