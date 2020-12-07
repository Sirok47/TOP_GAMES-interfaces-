package repository

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/model"
)

func Read (db *sql.DB,g *model.SingleGame,id int) {
	g= &model.SingleGame{0,"---",0,"---","---"}
	res, err := db.Query("select * from TopGames where id = $1",id)
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