package controller

import (
	"examples/repository"
	"examples/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	repo *repository.User
}

func NewUser() *User {
	ctr := new(User)
	ctr.repo = repository.NewUser()
	return ctr
}

func (ctr *User) Get(c *gin.Context) {
	user := ctr.repo.Get(1)
	var resp response.User
	resp.Mapping(user)

	c.JSON(http.StatusOK, resp)
}
