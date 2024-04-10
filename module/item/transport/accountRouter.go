package transport

import (
	controller "test3/module/item/controller"
)

func (t *Transport) SetAccountRoutes() {
	t.router.POST("/password", controller.PostPassword)
}
