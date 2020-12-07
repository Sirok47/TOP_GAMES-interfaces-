package repository

func (R TopGamesRepository) Delete(id int) {
	_,err := R.Db.Exec("delete from TopGames where id = $1", id)
	if err != nil{
		panic(err)
	}
}