package repository

import "fmt"

func (R TopGames) DeleteLine(id int) {
	_,err := R.Db.Exec("delete from TopGames where id = $1", id)
	if err != nil{
		fmt.Println(err)
	}
}