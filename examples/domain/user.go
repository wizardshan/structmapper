package domain

import "time"

type Users []*User

type User struct {
	ID int
	CreateTime time.Time
	UpdateTime *time.Time
	Mobile string
	Nickname string
	Money int
}

func (dom *User) Level() int {
	return 1
}


