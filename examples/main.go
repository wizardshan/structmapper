package main

import (
	"examples/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userCtr := controller.NewUser()
	usersCtr := controller.NewUsers()
	r.GET("user", userCtr.Get)
	r.GET("users", usersCtr.All)

	ordersCtr := controller.NewOrders()
	r.GET("orders", ordersCtr.All)

	r.Run()
}
