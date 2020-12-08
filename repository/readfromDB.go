package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
)

func (R TopGames) ReadLine(id int) (*model.SingleGame, error) {
	g := &model.SingleGame{Id: id, Name: "---", Rating: 0, Platform: "---", Date: "---"}
	res, err := R.db.Query("select * from TopGames where id = $1", id)
	defer res.Close()
	if err != nil {
		return g, err
	}
	for res.Next() {
		err = res.Scan(&g.Id, &g.Name, &g.Rating, &g.Platform, &g.Date)
		if err != nil {
			return g, err
		}
	}
	return g, err
}
