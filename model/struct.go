// Package model contains "SingleGame" structure
package model

// SingleGame stores data from requests
type SingleGame struct {
	ID       int     `json:"ID" bson:"_id"`
	Name     string  `json:"Name" bson:"Name"`
	Rating   float64 `json:"Rating" bson:"Rating"`
	Platform string  `json:"Platform" bson:"Platform"`
	Date     string  `json:"Date" bson:"Date"`
}
