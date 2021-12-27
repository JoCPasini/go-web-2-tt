package transacciones

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
	LastId() (int, error)
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

func (e *repository) LastId() (int, error) {
	return LastId, nil
}
