package domain

type Orders []*Order

type Order struct {
	ID int
	UserID int
	ShopID int
	Status int

	Consignee string
	Mobile string
	City string
	Province string

	Shop *Shop
	Items Items
}

func (dom Order) Expired() bool {
	return true
}
