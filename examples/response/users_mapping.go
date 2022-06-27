package response

import "examples/domain"

func (resp *Users) Mapping(dom domain.Users) {
	/**************** mapping start ****************/
	domUsersLen := len(dom)
	*resp = make(Users, domUsersLen)
	if domUsersLen > 0 {
		for domUsersIndex := 0; domUsersIndex < domUsersLen; domUsersIndex++ {
			domUsersItem := dom[domUsersIndex]
			respUser := new(User)
			respUser.ID = domUsersItem.ID
			respUser.CreateTime = DateTime(domUsersItem.CreateTime)
			if domUsersItem.UpdateTime != nil {
				updateTime := DateTime(*domUsersItem.UpdateTime)
				respUser.UpdateTime = &(updateTime)
			}

			//respUser.DeleteTime = fromStruct property not exist
			respUser.Mobile = domUsersItem.Mobile
			respUser.Nickname = domUsersItem.Nickname
			respUser.Money = Money(domUsersItem.Money)
			respUser.Level = domUsersItem.Level()
			(*resp)[domUsersIndex] = respUser
		}
	}

	/**************** mapping end  ****************/
}
