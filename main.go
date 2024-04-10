package main

import (
	"github.com/gin-gonic/gin"

	controller "test3/module/item/controller"
)

func main() {
	router := gin.Default()
	router.GET("/users", controller.GetUsers)
	router.GET("/users/:id", controller.GetSpecificUser)
	router.POST("/users", controller.PostUsers)
	router.DELETE("/users/:id", controller.DelUsers)
	router.PUT("/users/:id", controller.UpdateUser)
	router.Run("localhost:8080")

}
