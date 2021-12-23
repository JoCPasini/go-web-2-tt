package internal

var Transacciones = []Transaccion{}
var UltimoId int

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
	Store(id int, emisor string, receptor string) ([]Transaccion, error)
	//LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Transaccion, error) {
	return Transacciones, nil
}

func (e *repository) Store(id int, emisor string, receptor string) ([]Transaccion, error) {
	tr1 := Transaccion{id, "C0D1G0-D3-TR4NS4CC!0N", "pesos", 650.0, emisor, receptor, "20-20-2020"}
	UltimoId = id
	Transacciones = append(Transacciones, tr1)
	return Transacciones, nil
}
