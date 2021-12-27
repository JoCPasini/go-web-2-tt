package transacciones

type Service interface {
	GetAll() ([]Transaccion, error)
	Store(id int, emisor string, receptor string) (Transaccion, error)
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

func (s *service) Store(id int, emisor string, receptor string) (Transaccion, error) {

	LastId, err := s.repository.LastId()
	if err != nil {
		return Transaccion{}, err
	}
	LastId++

	transaccion, err := s.repository.Store(id, emisor, receptor)
	if err != nil {
		return Transaccion{}, err
	}
	return transaccion, nil
}
