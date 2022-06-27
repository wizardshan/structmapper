package response

type Orders []*Order

type Order struct {

	ID int `json:"id"`
	Status int `json:"status"`

	Consignee string `json:"consignee"`
	Mobile string `json:"mobile"`
	Province string `json:"province"`
	City string `json:"city"`

	Expired bool `json:"expired"`

	Shop *Shop `json:"shop"`
	Items []*Item `json:"items"`
}