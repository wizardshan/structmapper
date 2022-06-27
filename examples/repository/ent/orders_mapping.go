package ent

import "examples/domain"

func (o Orders) Mapping() domain.Orders {
	/**************** mapping start ****************/
	oOrdersLen := len(o)
	dom := make(domain.Orders, oOrdersLen)
	if oOrdersLen > 0 {
		for oOrdersIndex := 0; oOrdersIndex < oOrdersLen; oOrdersIndex++ {
			oOrdersItem := o[oOrdersIndex]
			domOrder := new(domain.Order)
			domOrder.ID = oOrdersItem.ID
			domOrder.UserID = oOrdersItem.UserID
			domOrder.ShopID = oOrdersItem.ShopID
			domOrder.Status = oOrdersItem.Status
			domOrder.Consignee = oOrdersItem.Consignee
			domOrder.Mobile = oOrdersItem.Mobile
			domOrder.City = oOrdersItem.City
			domOrder.Province = oOrdersItem.Province
			domOrder.Shop = oOrdersItem.Shop()
			domOrder.Items = oOrdersItem.Items()
			dom[oOrdersIndex] = domOrder
		}
		return dom
	}

	return nil
	/**************** mapping end  ****************/
}
