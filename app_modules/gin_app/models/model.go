package models

type Land struct {
	ID        int         `json:"id"`
	Acres     string      `json:"acres"`
	Buildings []*Building `json:"buildings"`
}

type Building struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Length int    `json:"length"`
	Height int    `json:"height"`
	Color  string `json:"color"`
}
