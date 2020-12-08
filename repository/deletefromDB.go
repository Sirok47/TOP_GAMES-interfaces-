package repository

func (R TopGames) DeleteLine(id int) error {
	_,err := R.Db.Exec("delete from TopGames where id = $1", id)
		return err
}