package transacciones

import "fmt"

var transacciones = []Transaccion{}
var LastId int

type Transaccion struct {
	Id                int     `json:"id"`
	CodigoTransaccion string  `json:"codigoTransaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fechaTransaccion`
}

type Repository interface {
	GetAll() ([]Transaccion, error)
	Store(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error)
	LastId(id int) (int, error)
	Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Transaccion, error) {
	return transacciones, nil
}

func (e *repository) Store(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error) {
	t1 := Transaccion{id, codigoTransaccion, moneda, monto, emisor, receptor, fechaTransaccion}
	transacciones = append(transacciones, t1)
	LastId = t1.Id
	return t1, nil
}

func (e *repository) LastId(id int) (int, error) {
	LastId = id
	return LastId, nil
}

func (e *repository) Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error) {
	t := Transaccion{CodigoTransaccion: codigoTransaccion, Moneda: moneda, Monto: monto, Emisor: emisor, Receptor: receptor, FechaTransaccion: fechaTransaccion}
	updated := false

	for i := range transacciones {
		if transacciones[i].Id == id {
			t.Id = id
			transacciones[i] = t
			updated = true
		}
	}
	if !updated {
		return Transaccion{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return t, nil

}
