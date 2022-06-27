package repository

import (
	"examples/domain"
	"examples/repository/ent"
	"time"
)

type User struct {
}

func NewUser() *User {
	return new(User)
}

func (repo *User) Get(id int) *domain.User {
	user := new(ent.User)
	user.ID = id
	user.Mobile = "1300000000"
	user.Nickname = "tom"
	user.Money = 101
	user.CreateTime = time.Now()
	return user.Mapping()
}

func (repo *User) All() domain.Users {
	user := new(ent.User)
	user.ID = 1
	user.Mobile = "1300000000"
	user.Nickname = "tom"
	user.Money = 101
	user.CreateTime = time.Now()
	var users ent.Users
	users = append(users, user)
	return users.Mapping()
}
