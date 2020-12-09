package repository

import "github.com/Sirok47/TOP_GAMES/model"

func (R TopGames) CreateLine(g *model.SingleGame) error {
	_, err := R.db.Exec("insert into TopGames (id,GameName,Rating,Platform,ReleaseDate) values ($1,$2,$3,$4,$5)", g.Id, g.Name, g.Rating, g.Platform, g.Date)
	return err
}
