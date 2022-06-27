package controller

import (
	"examples/repository"
	"examples/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Users struct {
	resp *repository.User
}

func NewUsers() *Users {
	ctr := new(Users)
	ctr.resp = new(repository.User)
	return ctr
}

func (ctr *Users) All(c *gin.Context) {
	users := ctr.resp.All()

	var resp response.Users
	resp.Mapping(users)

	c.JSON(http.StatusOK, resp)
}
