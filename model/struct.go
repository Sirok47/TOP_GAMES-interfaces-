package model

type SingleGame struct{
	Id int `json:"Id" bson:"_id"`
	Name string `json:"Name" bson:"Name"`
	Rating float64 `json:"Rating" bson:"Rating"`
	Platform string `json:"Platform" bson:"Platform"`
	Date string `json:"Date" bson:"Date"`
}

