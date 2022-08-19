package ticket

import "github.com/google/uuid"

type Repository interface {
	Create(ticket Ticket) (uuid.UUID, error)
	Get(id uuid.UUID) (Ticket, error)
	Delete(id uuid.UUID) error
}

type repository struct{}

func (r repository) Get(id uuid.UUID) (Ticket, error) {
	return Ticket{}, nil
}

func (r repository) Create(ticket Ticket) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}
func (r repository) Delete(id uuid.UUID) error {
	return nil
}

func NewRepository() Repository {
	return repository{}
}
