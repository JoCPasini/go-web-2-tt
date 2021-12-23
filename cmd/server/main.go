package main

import (
	"github.com/JosePasiniMercadolibre/cmd/server/handler"
	"github.com/JosePasiniMercadolibre/internal/transacciones"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := transacciones.NewRepository()
	service := transacciones.NewService(repo)
	p := handler.NewTransaccion(service)

	r := gin.Default()
	pr := r.Group("/transacciones")
	pr.POST("/")
}
