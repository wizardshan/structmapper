package repository

import (
	"examples/domain"
	"examples/repository/ent"
	"time"
)

type Order struct {
}

func NewOrder() *Order {
	return new(Order)
}

func (repo *Order) All() domain.Orders {
	order := new(ent.Order)
	order.ID = 1
	order.Status = 1
	order.Consignee = "tom"
	order.Mobile = "1300000000"
	order.Province = "ShangHai"
	order.City = "ShangHai"
	order.CreateTime = time.Now()

	shop := new(ent.Shop)
	shop.ID = 1
	shop.Name = "nike"
	order.Edges.Shop = shop

	item := new(ent.Item)
	item.ID = 1
	item.Title = "T shirt"
	item.Price = 14999
	order.Edges.Items = append(order.Edges.Items, item)

	var orders ent.Orders
	orders = append(orders, order)
	return orders.Mapping()
}
