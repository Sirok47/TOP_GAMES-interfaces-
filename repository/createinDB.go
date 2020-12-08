package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
)

func (R TopGames) CreateLine(g *model.SingleGame) string {
	_,err:=R.Db.Exec("insert into TopGames (id,GameName,Rating,Platform,ReleaseDate) values ($1,$2,$3,$4,$5)",g.Id,g.Name,g.Rating,g.Platform,g.Date)
	if err != nil{
		return "Can't create line."
	}
	str:= "Line have been created"
	return str
}