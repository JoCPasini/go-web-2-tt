package handler

import (
	"github.com/JosePasiniMercadolibre/go-web-2-tt/internal/transacciones"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id                int     `json:"id"`
	CodigoTransaccion string  `json:"codigoTransaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fechaTransaccion`
}

type Transaccion struct {
	service transacciones.Service
}

func NewTransaccion(t transacciones.Service) *Transaccion {
	return &Transaccion{
		service: t,
	}
}

func (t *Transaccion) GetAll() gin.HandlerFunc {
	return func(ctx gin.Context) {
		tokem := ctx.Request.Header.Get("tokem")
		if tokem != "123456" {
			ctx.JSON(400, gin.H{
				"error": "Tokem Inválido",
			})
			return
		}

		transacciones, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": "No se encontró",
			})
			return
		}
		ctx(200, transacciones)
	}
}
