package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (r TopGames) Create(g *model.SingleGame) error {
	_, err := r.db.InsertOne(
		r.ctx,
		bson.M{
			"_id":      g.ID,
			"Name":     g.Name,
			"Rating":   g.Rating,
			"Platform": g.Platform,
			"Date":     g.Date,
		})

	return errors.Wrap(err, "Create failed")
}
