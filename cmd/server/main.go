package main

import (
	"github.com/JosePasiniMercadolibre/go-web-2-tt/cmd/server/handler"
	"github.com/JosePasiniMercadolibre/go-web-2-tt/internal/transacciones"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := transacciones.NewRepository()
	service := transacciones.NewService(repo)
	handler := handler.NewTransaccion(service)

	r := gin.Default()
	pr := r.Group("/transacciones")
	pr.GET("/", handler.GetAll())

	r.Run()
}
