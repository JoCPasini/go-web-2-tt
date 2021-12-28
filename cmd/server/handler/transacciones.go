package handler

import (
	"fmt"
	"os"
	"strconv"

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
		/*
			tokem := ctx.Request.Header.Get("token")
			if tokem != os.Getenv("TOKEN") {
				ctx.JSON(400, gin.H{
					"error": "Token Inválido",
				})
				return
			}
		*/

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
		/*
			tokem := ctx.Request.Header.Get("token")
			if tokem != os.Getenv("TOKEN") {
				ctx.JSON(400, gin.H{
					"error": "Token Inválido",
				})
				return
			}
		*/
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
func (t *Transaccion) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/*
			tokem := ctx.Request.Header.Get("token")
			if tokem != os.Getenv("TOKEN") {
				ctx.JSON(400, gin.H{
					"error": "Token Inválido",
				})
				return
			}
		*/
		idParam, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "id invalido"})
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Condicional para verificar que el campo no venga vacío
		if req.CodigoTransaccion == "" {
			ctx.JSON(400, gin.H{"error": "El codigo de la transaccion es requerido, no puede ir un campo vacío, debe modificar todo"})
			return
		}

		if req.Moneda == "" {
			ctx.JSON(400, gin.H{"error": "La moneda de la transaccion es requerida, no puede ir un campo vacío, debe modificar todo"})
			return
		}

		if req.Monto == 0 {
			ctx.JSON(400, gin.H{"error": "El monto de la transaccion es requerido, no puede ir un campo vacío, debe modificar todo"})
			return
		}

		if req.Emisor == "" {
			ctx.JSON(400, gin.H{"error": "El emisor de la transaccion es requerido, no puede ir un campo vacío, debe modificar todo"})
			return
		}

		if req.Receptor == "" {
			ctx.JSON(400, gin.H{"error": "El receptor de la transaccion es requerido, no puede ir un campo vacío, debe modificar todo"})
			return
		}

		if req.FechaTransaccion == "" {
			ctx.JSON(400, gin.H{"error": "La fecha de la transaccion es requerido, no puede ir un campo vacío, debe modificar todo"})
			return
		}

		t1, err := t.service.Update(int(idParam), req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, t1)

	}
}

func (t *Transaccion) UpdateCodigoMonto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokem := ctx.Request.Header.Get("token")
		if tokem != os.Getenv("TOKEN") {
			ctx.JSON(400, gin.H{
				"error": "Token Inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "id invalido"})
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.CodigoTransaccion == "" {
			ctx.JSON(400, gin.H{"error": "El codigo de la transaccion es requerido, no puede ir un campo vacío, debe modificar todo"})
			return
		}

		if req.Monto == 0 {
			ctx.JSON(400, gin.H{"error": "El monto de la transaccion es requerido, no puede ir un campo vacío, debe modificar todo"})
			return
		}

		t1, err := t.service.UpdateCodigoMonto(int(id), req.CodigoTransaccion, req.Monto)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id no encontrado"})
		}

		ctx.JSON(200, t1)

	}
}

func (t *Transaccion) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokem := ctx.Request.Header.Get("token")
		if tokem != os.Getenv("TOKEN") {
			ctx.JSON(400, gin.H{
				"error": "Token Inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{
			"data": fmt.Sprintf("La transacción %d ha sido eliminada", id)})
	}
}
