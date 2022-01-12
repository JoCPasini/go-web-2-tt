package transacciones

import (
	"encoding/json"
	"testing"

	"github.com/JosePasiniMercadolibre/go-web-2-tt/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAllServices(t *testing.T) {
	input := []Transaccion{
		{
			Id:                1,
			CodigoTransaccion: "1",
			Moneda:            "1",
			Monto:             1,
			Emisor:            "1",
			Receptor:          "1",
			FechaTransaccion:  "1",
		},
		{
			Id:                2,
			CodigoTransaccion: "2",
			Moneda:            "2",
			Monto:             2,
			Emisor:            "2",
			Receptor:          "2",
			FechaTransaccion:  "2",
		},
	}

	inputJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: inputJson,
		Err:  nil,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	resultado, err := myService.GetAll()
	assert.Equal(t, input, resultado)
	assert.Nil(t, err)
}

/*
func TestUpdateService(t *testing.T) {

	input := Transaccion{
		Id:                1,
		CodigoTransaccion: "1",
		Moneda:            "1",
		Monto:             1,
		Emisor:            "1",
		Receptor:          "1",
		FechaTransaccion:  "1",
	}
	inputJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: inputJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	resultado, _ := myService.Store(1, "CodigoTransaccion modificado", "Moneda modificado", 99, "Emisor modificado", "Receptor modificado", "FechaTransaccion modificado")
	//resultado, _ := myService.Update(1, "CodigoTransaccion modificado", "Moneda modificado", 99, "Emisor modificado", "Receptor modificado", "FechaTransaccion modificado")
	assert.Equal(t, input, resultado)

	//assert.Nil(t, err)
}
*/
