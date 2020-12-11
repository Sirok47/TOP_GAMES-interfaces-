package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Read gets data from DB by ID and writes it into "TopGames"'s object
func (r *TopGames) Read(id int) (*model.SingleGame, error) {
	g := &model.SingleGame{ID: id, Name: "---", Rating: 0, Platform: "---", Date: "---"}
	err := r.db.FindOne(
		r.ctx,
		bson.D{
			primitive.E{
				Key:   "_id",
				Value: g.ID,
			},
		}).Decode(&g)

	return g, errors.Wrap(err, "Read failed")
}
