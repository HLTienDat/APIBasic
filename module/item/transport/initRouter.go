package transport

import (
	"github.com/gin-gonic/gin"
)

type Transport struct {
	router *gin.Engine
}

func NewTransport() *Transport {
	return &Transport{
		router: gin.Default(),
	}
}

func (t *Transport) Run() {
	t.router.Run("localhost:8080")
}
