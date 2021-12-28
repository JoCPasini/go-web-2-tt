package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/JosePasiniMercadolibre/go-web-2-tt/internal/transacciones"
	"github.com/JosePasiniMercadolibre/go-web-2-tt/pkg/web"
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

// ListTransacciones godoc
// @Sumary List transacciones
// @Tags Transacciones
// @Description get transacciones
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transacciones/getAll [get]
func (t *Transaccion) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokem := ctx.Request.Header.Get("token")
		if tokem != os.Getenv("TOKEN") {
			ctx.JSON(400, web.NewResponse(400, nil, "Token inválido"))
			return
		}

		transacciones, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, string(err.Error())))
			return
		}
		ctx.JSON(200, web.NewResponse(200, transacciones, ""))
	}
}

func (t *Transaccion) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokem := ctx.Request.Header.Get("token")
		if tokem != os.Getenv("TOKEN") {
			ctx.JSON(400, web.NewResponse(400, nil, "Token inválido"))
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
func (t *Transaccion) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokem := ctx.Request.Header.Get("token")
		if tokem != os.Getenv("TOKEN") {
			ctx.JSON(400, web.NewResponse(400, nil, "Token inválido"))
			return
		}
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
			ctx.JSON(400, web.NewResponse(400, nil, "Error, el código de transacción no puede ir vacío"))
			return
		}

		if req.Moneda == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Error, la moneda no puede ir vacía"))
			return
		}

		if req.Monto == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Error, el monto no puede ir vacío"))
			return
		}

		if req.Emisor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Error, el emisor no puede ir vacío"))
			return
		}

		if req.Receptor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Error, el receptor no puede ir vacío"))
			return
		}

		if req.FechaTransaccion == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Error, la fecha de transacción no puede ir vacío"))
			return
		}

		t1, err := t.service.Update(int(idParam), req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, string(err.Error())))
			return
		}
		ctx.JSON(200, t1)

	}
}

func (t *Transaccion) UpdateCodigoMonto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokem := ctx.Request.Header.Get("token")
		if tokem != os.Getenv("TOKEN") {
			ctx.JSON(400, web.NewResponse(400, nil, "Token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id inválido"))

			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, string(err.Error())))
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

		ctx.JSON(200, web.NewResponse(200, t1, ""))

	}
}

func (t *Transaccion) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokem := ctx.Request.Header.Get("token")
		if tokem != os.Getenv("TOKEN") {
			ctx.JSON(400, web.NewResponse(400, nil, "Token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, string(err.Error())))
			return
		}
		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, string(err.Error())))
			return
		}

		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("Transacción %d eliminada", id), ""))
	}
}
