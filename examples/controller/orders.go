package controller

import (
	"examples/repository"
	"examples/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Orders struct {
	resp *repository.Order
}

func NewOrders() *Orders {
	ctr := new(Orders)
	ctr.resp = new(repository.Order)
	return ctr
}

func (ctr *Orders) All(c *gin.Context) {
	orders := ctr.resp.All()

	var resp response.Orders
	resp.Mapping(orders)

	c.JSON(http.StatusOK, resp)
}
