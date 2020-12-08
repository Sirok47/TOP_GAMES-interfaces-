package repository

func (R TopGames) DeleteLine(id int) error {
	_, err := R.db.Exec("delete from TopGames where id = $1", id)
	return err
}
