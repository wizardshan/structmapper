package ent

import "examples/domain"

func (u Users) Mapping() domain.Users {
	/**************** mapping start ****************/
	uUsersLen := len(u)
	dom := make(domain.Users, uUsersLen)
	if uUsersLen > 0 {
		for uUsersIndex := 0; uUsersIndex < uUsersLen; uUsersIndex++ {
			uUsersItem := u[uUsersIndex]
			domUser := new(domain.User)
			domUser.ID = uUsersItem.ID
			domUser.CreateTime = uUsersItem.CreateTime
			domUser.UpdateTime = &uUsersItem.UpdateTime
			domUser.Mobile = uUsersItem.Mobile
			domUser.Nickname = uUsersItem.Nickname
			domUser.Money = uUsersItem.Money
			dom[uUsersIndex] = domUser
		}
		return dom
	}

	return nil
	/**************** mapping end  ****************/
}
