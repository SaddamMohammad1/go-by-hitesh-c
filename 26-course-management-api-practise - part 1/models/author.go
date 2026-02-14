package models

type Author struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}
