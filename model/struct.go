package model

type SingleGame struct{
	Id int `json:"Id"`
	Name string `json:"Name"`
	Rating float64 `json:"Rating"`
	Platform string `json:"Platform"`
	Date string `json:"Date"`
}
type JSON struct{
	Id string `json:"Id"`
	Name string `json:"Name"`
	Rating string `json:"Rating"`
	Platform string `json:"Platform"`
	Date string `json:"Date"`
}

