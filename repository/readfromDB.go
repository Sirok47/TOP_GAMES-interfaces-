package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
)

func (R TopGamesRepository) Read (g *model.SingleGame,id int) {
	g= &model.SingleGame{0,"---",0,"---","---"}
	res, err := R.Db.Query("select * from TopGames where id = $1",id)
	if err != nil {
		panic(err)
	}
	for res.Next(){
		err =res.Scan(&g.Id,&g.Name, &g.Rating, &g.Platform, &g.Date)
		if err != nil {
			panic(err)
		}
	}
	res.Close()
}