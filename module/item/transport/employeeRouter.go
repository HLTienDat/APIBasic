package transport

import (
	controller "test3/module/item/controller"
)

func (t *Transport) SetEmployeeRoutes() {
	t.router.GET("/users", controller.GetUsers)
	t.router.GET("/users/:id", controller.GetSpecificUser)
	t.router.POST("/users", controller.PostUsers)
	t.router.DELETE("/users/:id", controller.DelUsers)
	t.router.PUT("/users/:id", controller.UpdateUser)
}
