package response

import "examples/domain"

func (resp *User) Mapping(dom *domain.User) {
	/**************** mapping start ****************/
	resp.ID = dom.ID
	resp.CreateTime = DateTime(dom.CreateTime)
	if dom.UpdateTime != nil {
		updateTime := DateTime(*dom.UpdateTime)
		resp.UpdateTime = &(updateTime)
	}

	//resp.DeleteTime = fromStruct property not exist
	resp.Mobile = dom.Mobile
	resp.Nickname = dom.Nickname
	resp.Money = Money(dom.Money)
	resp.Level = dom.Level()

	/**************** mapping end  ****************/
}
