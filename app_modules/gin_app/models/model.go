package models

type Land struct {
	ID        int         `json:"id"`
	Acres     string      `json:"acres"`
	Buildings []*Building `json:"buildings"`
}

type Building struct {
	Name string `json:"name"`
}
