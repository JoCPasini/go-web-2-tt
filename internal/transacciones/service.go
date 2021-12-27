package transacciones

type Service interface {
	GetAll() ([]Transaccion, error)
	Store(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]Transaccion, error) {
	transacciones, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return transacciones, nil
}

func (s *service) Store(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error) {

	LastId, err := s.repository.LastId()
	if err != nil {
		return Transaccion{}, err
	}
	LastId++

	transaccion, err := s.repository.Store(id, codigoTransaccion, moneda, monto, emisor, receptor, fechaTransaccion)
	if err != nil {
		return Transaccion{}, err
	}
	return transaccion, nil
}
