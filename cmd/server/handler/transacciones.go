package handler

import (
	"github.com/JosePasiniMercadolibre/go-web-2-tt/internal/transacciones"
	"github.com/gin-gonic/gin"
)

type Transaccion struct {
	service transacciones.Service
}

type request struct {
	Id                int     `json:"id"`
	CodigoTransaccion string  `json:"codigoTransaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fechaTransaccion`
}

func NewTransaccion(t transacciones.Service) *Transaccion {
	return &Transaccion{
		service: t,
	}
}

func (t *Transaccion) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
		ctx.JSON(200, transacciones)
	}
}

func (t *Transaccion) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokem := ctx.Request.Header.Get("tokem")
		if tokem != "123456" {
			ctx.JSON(400, gin.H{
				"error": "Tokem Inválido",
			})
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		t, err := t.service.Store(req.Id, req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, t)
	}
}
