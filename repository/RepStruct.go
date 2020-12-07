package repository

import "database/sql"

type TopGamesRepository struct {
	Db *sql.DB
}
