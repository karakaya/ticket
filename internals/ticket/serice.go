package ticket

import "github.com/google/uuid"

type Service interface {
	CreateTicket(ticket Ticket) (uuid.UUID, error)
	GetTicket(id uuid.UUID) (Ticket, error)
	DeleteTicket(id uuid.UUID) error
}

type service struct {
	repo Repository
}

func (s service) CreateTicket(ticket Ticket) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (s service) GetTicket(id uuid.UUID) (Ticket, error) {
	return Ticket{}, nil
}

func (s service) DeleteTicket(id uuid.UUID) error {
	return nil
}

func NewService(r Repository) Service {
	return service{repo: r}
}
