package ticket

import (
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateTicket(ticketRequest CreateTicketRequest) (uuid.UUID, error)
	GetTicket(id uuid.UUID) (Ticket, error)
	DeleteTicket(id uuid.UUID) error
}

type service struct {
	repo Repository
}

func (s service) CreateTicket(ticketRequest CreateTicketRequest) (uuid.UUID, error) {
	var ticket Ticket

	ticket.ID = uuid.New()
	ticket.Email = ticketRequest.Email
	ticket.Title = ticketRequest.Title
	ticket.Body = ticketRequest.Body
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()
	return s.repo.Create(ticket)
}

func (s service) GetTicket(id uuid.UUID) (Ticket, error) {
	return s.repo.Get(id)
}

func (s service) DeleteTicket(id uuid.UUID) error {
	return nil
}

func NewService(r Repository) Service {
	return service{repo: r}
}
