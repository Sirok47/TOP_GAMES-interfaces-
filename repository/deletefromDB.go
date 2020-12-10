package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (R TopGames) DeleteLine(id int) error {
	_, err := R.db.DeleteOne(R.ctx,bson.D{primitive.E{Key: "_id", Value: id}})
	return err
}
