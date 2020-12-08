package repository

func (R TopGames) DeleteLine(id int) string {
	_,err := R.Db.Exec("delete from TopGames where id = $1", id)
	if err != nil{
		return "Can't delete line."
	}
	return "Line have been deleted."
}