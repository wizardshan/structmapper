package domain

import "time"

type Items []*Item

type Item struct {
	ID int `json:"id"`
	CreateTime time.Time
	UpdateTime time.Time
	OrderID int
	Title string

	Price int
}
