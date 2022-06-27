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