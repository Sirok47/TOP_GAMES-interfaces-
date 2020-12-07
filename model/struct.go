package model

import "database/sql"

type TopGamesHandler struct {
	Db *sql.DB
	G SingleGame
}
type SingleGame struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Rating float64 `json:"reit"`
	Platform string `json:"platform"`
	Date string `json:"date"`
}
