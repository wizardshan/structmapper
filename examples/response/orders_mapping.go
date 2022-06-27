package response

import "examples/domain"

func (resp *Orders) Mapping(dom domain.Orders) {
	/**************** mapping start ****************/
	domOrdersLen := len(dom)
	*resp = make(Orders, domOrdersLen)
	if domOrdersLen > 0 {
		for domOrdersIndex := 0; domOrdersIndex < domOrdersLen; domOrdersIndex++ {
			domOrdersItem := dom[domOrdersIndex]
			respOrder := new(Order)
			respOrder.ID = domOrdersItem.ID
			respOrder.Status = domOrdersItem.Status
			respOrder.Consignee = domOrdersItem.Consignee
			respOrder.Mobile = domOrdersItem.Mobile
			respOrder.Province = domOrdersItem.Province
			respOrder.City = domOrdersItem.City
			respOrder.Expired = domOrdersItem.Expired()

			if domOrdersItem.Shop != nil {
				respOrderShop := new(Shop)
				respOrderShop.ID = domOrdersItem.Shop.ID
				respOrderShop.Name = domOrdersItem.Shop.Name
				respOrder.Shop = respOrderShop
			}

			domOrdersItemItemsLen := len(domOrdersItem.Items)
			respOrderItems := make([]*Item, domOrdersItemItemsLen)
			if domOrdersItemItemsLen > 0 {
				for domOrdersItemItemsIndex := 0; domOrdersItemItemsIndex < domOrdersItemItemsLen; domOrdersItemItemsIndex++ {
					domOrdersItemItemsItem := domOrdersItem.Items[domOrdersItemItemsIndex]
					respOrderItem := new(Item)
					respOrderItem.ID = domOrdersItemItemsItem.ID
					respOrderItem.Title = domOrdersItemItemsItem.Title
					respOrderItem.Price = Money(domOrdersItemItemsItem.Price)
					respOrderItems[domOrdersItemItemsIndex] = respOrderItem
				}
			}

			respOrder.Items = respOrderItems
			(*resp)[domOrdersIndex] = respOrder
		}
	}

	/**************** mapping end  ****************/
}
