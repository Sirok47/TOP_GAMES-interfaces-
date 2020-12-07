package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
)

func (R TopGamesRepository) Write () {
	_,err:=R.Db.Exec("insert into TopGames (id,GameName,Rating,Platform,ReleaseDate) values ($1,$2,$3,$4,$5)",model.G.Id,model.G.Name,model.G.Rating,model.G.Platform,model.G.Date)
	if err != nil{
		panic(err)
	}
}