package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
)

func (R TopGames) ReadLine (id int) *model.SingleGame {
	g:= &model.SingleGame{Id:id,Name: "---",Rating: 0,Platform: "---",Date: "---"}
	res, err := R.Db.Query("select * from TopGames where id = $1",id)
	defer res.Close()
	if err != nil {
		return nil
	}
	for res.Next(){
		err =res.Scan(&g.Id,&g.Name, &g.Rating, &g.Platform, &g.Date)
		if err != nil {
			return nil
		}
	}
	return g
}