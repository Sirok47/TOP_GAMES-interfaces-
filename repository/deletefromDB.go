package repository

import (
	"database/sql"
)

func Delete(db *sql.DB,id int) {
	_,err := db.Exec("delete from TopGames where id = $1", id)
	if err != nil{
		panic(err)
	}
}