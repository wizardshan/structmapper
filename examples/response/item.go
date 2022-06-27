package response

type Item struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Price Money `json:"price"`
}
