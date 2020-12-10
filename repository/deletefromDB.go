package repository

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r TopGames) Delete(id int) error {
	_, err := r.db.DeleteOne(
		r.ctx,
		bson.D{
			primitive.E{
				Key:   "_id",
				Value: id,
			},
		})

	return errors.Wrap(err, "Delete failed")
}
