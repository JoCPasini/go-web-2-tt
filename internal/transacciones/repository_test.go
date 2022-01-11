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
	req1 := Transaccion{1, "1", "1", 1, "1", "1", "1"}
	req2 := Transaccion{2, "2", "2", 2, "2", "2", "2"}
	dataAux, _ := json.Marshal([]Transaccion{req1, req2})
	json.Unmarshal(dataAux, &data)
	s.readExecute = true
	return nil
}

func (s *stubStore) Read(data interface{}) error {
	req1 := Transaccion{1, "1", "1", 1, "1", "1", "1"}
	req2 := Transaccion{2, "2", "2", 2, "2", "2", "2"}

	dataAux, _ := json.Marshal([]Transaccion{req1, req2})
	json.Unmarshal(dataAux, &data)
	return nil
}

func (s *stubStore) Write(data interface{}) error { return nil }
func (s *mockStore) Write(data interface{}) error { return nil }

func TestGetAll(t *testing.T) {
	db := stubStore{}
	repo := NewRepository(&db)
	req1 := Transaccion{1, "1", "1", 1, "1", "1", "1"}
	req2 := Transaccion{2, "2", "2", 2, "2", "2", "2"}
	data2 := []Transaccion{req1, req2}
	respuesta, err := repo.GetAll()
	assert.Equal(t, respuesta, data2, err)
}

func TestUpdate(t *testing.T) {
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
