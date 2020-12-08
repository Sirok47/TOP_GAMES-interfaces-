package repository

import "database/sql"

type TopGames struct {
	db *sql.DB
}
func NewRep(db *sql.DB) *TopGames {
	return &TopGames{db}
}