package transacciones

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubStore struct{}
type mockStore struct {
	readExecute bool
}

func (s *mockStore) Read(data interface{}) error {

	transacciones := []Transaccion{
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
		{
			Id:                3,
			CodigoTransaccion: "3",
			Moneda:            "3",
			Monto:             3,
			Emisor:            "3",
			Receptor:          "3",
			FechaTransaccion:  "3",
		},
	}

	dataAux, _ := json.Marshal(transacciones)
	err := json.Unmarshal(dataAux, &data)
	if err != nil {
		return err
	}
	s.readExecute = true
	return nil
}

func (s *stubStore) Read(data interface{}) error {
	transacciones := []Transaccion{
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
		{
			Id:                3,
			CodigoTransaccion: "3",
			Moneda:            "3",
			Monto:             3,
			Emisor:            "3",
			Receptor:          "3",
			FechaTransaccion:  "3",
		},
	}

	dataAux, _ := json.Marshal(transacciones)
	err := json.Unmarshal(dataAux, &data)
	if err != nil {
		return err
	}
	return nil
}

func (s *stubStore) Write(data interface{}) error { return nil }
func (s *mockStore) Write(data interface{}) error { return nil }

func TestGetAllRepository(t *testing.T) {
	db := stubStore{}
	repo := NewRepository(&db)
	// Creo una lista de objetos identicos a los creados arriba en el Mock (para comparar)
	req1 := Transaccion{1, "1", "1", 1, "1", "1", "1"}
	req2 := Transaccion{2, "2", "2", 2, "2", "2", "2"}
	req3 := Transaccion{3, "3", "3", 3, "3", "3", "3"}
	data := []Transaccion{req1, req2, req3}

	respuesta, err := repo.GetAll()

	assert.Equal(t, respuesta, data, err)
	assert.NotNil(t, respuesta)
	assert.NoError(t, err)
}

func TestUpdateCodigoMontoRepository(t *testing.T) {
	type TransaccionCambio struct {
		CodigoTransaccion string
		Monto             float64
	}

	db := mockStore{}
	repo := NewRepository(&db)

	esperado := Transaccion{1, "After Transaction", "1", 999, "1", "1", "1"}
	cambios := TransaccionCambio{"After Transaction", 999}

	resultado, err := repo.UpdateCodigoMonto(1, cambios.CodigoTransaccion, cambios.Monto)
	assert.Nil(t, err, "Hubo un error")
	assert.Equal(t, esperado, resultado, err)
	assert.Equal(t, true, db.readExecute, "No se ejecutó el Read")
}

func TestStoreRepository(t *testing.T) {
	t1 := Transaccion{
		Id:                5,
		CodigoTransaccion: "5",
		Moneda:            "5",
		Monto:             5,
		Emisor:            "5",
		Receptor:          "5",
		FechaTransaccion:  "5",
	}
	db := mockStore{}
	repo := NewRepository(&db)
	resultado, _ := repo.Store(t1.Id, t1.CodigoTransaccion, t1.Moneda, t1.Monto, t1.Emisor, t1.Receptor, t1.FechaTransaccion)
	assert.Equal(t, resultado.Id, 5)
	assert.NotNil(t, resultado)
}

func TestDeleteRepository(t *testing.T) {
	db := mockStore{}
	repo := NewRepository(&db)

	err := repo.Delete(1)
	assert.Nil(t, err)
}

func TestUpdateRepository(t *testing.T) {
	t1 := Transaccion{
		Id:                1,
		CodigoTransaccion: "4 - Updateada",
		Moneda:            "4 - Updateada",
		Monto:             4,
		Emisor:            "4 - Updateada",
		Receptor:          "4 - Updateada",
		FechaTransaccion:  "4 - Updateada",
	}
	db := mockStore{}
	repo := NewRepository(&db)

	t1, err := repo.Update(t1.Id, t1.CodigoTransaccion, t1.Moneda, t1.Monto, t1.Emisor, t1.Receptor, t1.FechaTransaccion)

	assert.Equal(t, t1.CodigoTransaccion, "4 - Updateada")
	assert.Equal(t, t1.Moneda, "4 - Updateada")
	assert.Equal(t, t1.Monto, 4.0)
	assert.Equal(t, t1.Emisor, "4 - Updateada")
	assert.Equal(t, t1.Receptor, "4 - Updateada")
	assert.Equal(t, t1.FechaTransaccion, "4 - Updateada")
	assert.Nil(t, err)
}

func TestLasIdRepository(t *testing.T) {

	db := mockStore{}
	repo := NewRepository(&db)

	resultado, _ := repo.LastId()

	// El 3 es el último ID de la "base de datos (Mock)" arriba se puede ver linea:19, func:Read()
	assert.Equal(t, resultado, 3)

}
