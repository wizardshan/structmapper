package ent

import "examples/domain"

func (i Items) Mapping() domain.Items {
	/**************** mapping start ****************/
	iItemsLen := len(i)
	if iItemsLen > 0 {
		dom := make(domain.Items, iItemsLen)
		for iItemsIndex := 0; iItemsIndex < iItemsLen; iItemsIndex++ {
			iItemsItem := i[iItemsIndex]
			domItem := new(domain.Item)
			domItem.ID = iItemsItem.ID
			domItem.CreateTime = iItemsItem.CreateTime
			domItem.UpdateTime = iItemsItem.UpdateTime
			domItem.OrderID = iItemsItem.OrderID
			domItem.Title = iItemsItem.Title
			domItem.Price = iItemsItem.Price
			dom[iItemsIndex] = domItem
		}
		return dom
	}

	return nil
	/**************** mapping end  ****************/
}
