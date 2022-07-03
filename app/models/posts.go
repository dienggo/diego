package models

type Posts []struct {
	Body   string `json:"body"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	UserID int64  `json:"userId"`
}