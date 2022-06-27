# structmapper 

structmapper is tool to auto generate struct to struct mapping function

```sh
go get -u github.com/wizardshan/structmapper
```

## Features
* mapper from struct to struct with same name
* mapper from slice to slice
* mapper from method to field with same name

```go
//go:generate go run github.com/wizardshan/structmapper -toName User -fromName User -toPath ./response  -fromPath ./domain -toVar resp -fromVar dom
//go:generate go run github.com/wizardshan/structmapper -toName Orders -fromName Orders -toPath ./response  -fromPath ./domain -toVar resp -fromVar dom


package response

type Users []*User

type User struct {
	ID int `json:"id"`
	CreateTime DateTime `json:"createTime"`
	UpdateTime *DateTime `json:"updateTime"`
	DeleteTime *DateTime `json:"deleteTime"`
	Mobile string `json:"mobile"`
	Nickname string `json:"nickname"`
	Money Money `json:"money"`
	Level int `json:"level"`
}


package domain

type Users []*User

type User struct {
	ID int
	CreateTime time.Time
	UpdateTime *time.Time
	Mobile string
	Nickname string
	Money int
}

func (dom *User) Level() int {
	return 1
}


user_mapping.go

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


orders_mapping.go

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

```


## Benchmarks

| Benchmark name                 |       (1) |             (2) |          (3) |             (4) |
| ------------------------------ | ---------:| ---------------:| ------------:| ---------------:|
| BenchmarkMapper         | **1000000000** | **0.3679 ns/op** |   **0 B/op** | **0 allocs/op** |
| BenchmarkCopier         |     163282 |     6852 ns/op |       3000 B/op  |     48 allocs/op |
