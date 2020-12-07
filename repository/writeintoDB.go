package repository

import (
	"database/sql"
	"github.com/Sirok47/TOP_GAMES/model"
)

func Write (db *sql.DB) {
	_,err:=db.Exec("insert into TopGames (id,GameName,Rating,Platform,ReleaseDate) values ($1,$2,$3,$4,$5)",model.G.Id,model.G.Name,model.G.Rating,model.G.Platform,model.G.Date)
	if err != nil{
		panic(err)
	}
	return
}