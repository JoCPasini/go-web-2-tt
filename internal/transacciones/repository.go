package transacciones

import (
	"fmt"

	"github.com/JosePasiniMercadolibre/go-web-2-tt/pkg/store"
)

/*
var transacciones = []Transaccion{}
var LastId int
*/

//var trans []Transaccion

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
	LastId() (int, error)
	Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error)
	UpdateCodigoMonto(id int, codigoTransaccion string, monto float64) (Transaccion, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Transaccion, error) {
	var trans []Transaccion
	r.db.Read(&trans)
	return trans, nil
}

func (r *repository) Store(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error) {

	var trans []Transaccion
	r.db.Read(&trans)
	t1 := Transaccion{id, codigoTransaccion, moneda, monto, emisor, receptor, fechaTransaccion}
	trans = append(trans, t1)
	if err := r.db.Write(trans); err != nil {
		return Transaccion{}, err
	}
	return t1, nil

}

func (r *repository) LastId() (int, error) {
	var trans []Transaccion
	if err := r.db.Read(&trans); err != nil {
		return 0, err
	}
	if len(trans) == 0 {
		return 0, nil
	}
	return trans[len(trans)-1].Id, nil
}

func (r *repository) Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error) {

	var trans []Transaccion

	t := Transaccion{CodigoTransaccion: codigoTransaccion, Moneda: moneda, Monto: monto, Emisor: emisor, Receptor: receptor, FechaTransaccion: fechaTransaccion}
	r.db.Read(&trans)
	updated := false

	for i := range trans {
		if trans[i].Id == id {
			t.Id = id
			trans[i] = t
			updated = true
		}
	}
	if !updated {
		return Transaccion{}, fmt.Errorf("Transaccion %d no encontrado", id)
	}
	//trans[id] = t
	if err := r.db.Write(trans); err != nil {
		return Transaccion{}, err
	}
	return t, nil
}

func (r *repository) UpdateCodigoMonto(id int, codigoTransaccion string, monto float64) (Transaccion, error) {
	var t1 Transaccion
	var trans []Transaccion
	r.db.Read(&trans)
	updated := false

	for i := range trans {
		if trans[i].Id == id {
			trans[i].CodigoTransaccion = codigoTransaccion
			trans[i].Monto = monto
			updated = true
			t1 = trans[i]
		}
	}
	if !updated {
		return Transaccion{}, fmt.Errorf("Transaccion %d no encontrado", id)
	}
	if err := r.db.Write(trans); err != nil {
		return Transaccion{}, err
	}
	return t1, nil
}

func (r *repository) Delete(id int) error {
	var trans []Transaccion
	r.db.Read(&trans)
	deleted := false
	var index int
	for i := range trans {
		if trans[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Transaccion %d no encontrada", id)
	}

	trans = append(trans[:index], trans[index+1:]...)
	if err := r.db.Write(trans); err != nil {
		return err
	}
	return nil
}
