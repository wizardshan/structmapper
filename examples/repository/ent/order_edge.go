package ent

import (
	"examples/domain"
)

// manually

func (o *Order) Shop() *domain.Shop {
	if o.Edges.Shop == nil {
		return nil
	}

	s := o.Edges.Shop
	dom := new(domain.Shop)
	dom.ID = s.ID
	dom.Name = s.Name
	return dom
}

func (o *Order) Items() domain.Items {
	i := o.Edges.Items

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
}

