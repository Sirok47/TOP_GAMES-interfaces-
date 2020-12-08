package repository

import "github.com/Sirok47/TOP_GAMES/model"

func (R TopGames) UpdateLine(g *model.SingleGame) error {
	_, err := R.Db.Exec("UPDATE TopGames SET GameName = $1, Rating = $2, Platform = $3, ReleaseDate = $4 WHERE Id = $5", g.Name, g.Rating, g.Platform, g.Date, g.Id)
	if err != nil {
		return err
	}
	return err
}