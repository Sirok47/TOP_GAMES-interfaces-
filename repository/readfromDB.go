package repository

import (
	"fmt"
	"github.com/Sirok47/TOP_GAMES/model"
)

func (R TopGames) ReadLine (id int) *model.SingleGame {
	g:= &model.SingleGame{Id:id,Name: "---",Rating: 0,Platform: "---",Date: "---"}
	res, err := R.Db.Query("select * from TopGames where id = $1",id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next(){
		err =res.Scan(&g.Id,&g.Name, &g.Rating, &g.Platform, &g.Date)
		if err != nil {
			fmt.Println(err)
		}
	}
	res.Close()
	return g
}