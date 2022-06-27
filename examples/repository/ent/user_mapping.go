package ent

import "examples/domain"

func (u *User) Mapping() *domain.User {
	/**************** mapping start ****************/
	dom := new(domain.User)
	dom.ID = u.ID
	dom.CreateTime = u.CreateTime
	dom.UpdateTime = &u.UpdateTime
	dom.Mobile = u.Mobile
	dom.Nickname = u.Nickname
	dom.Money = u.Money
	return dom

	/**************** mapping end  ****************/
}
