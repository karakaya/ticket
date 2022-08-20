package ticket

import (
	"time"

	"ticket/internals/rabbit"

	"github.com/google/uuid"
)

type Service interface {
	CreateTicket(ticketRequest CreateTicketRequest) (uuid.UUID, error)
	GetTicket(id uuid.UUID) (Ticket, error)
	DeleteTicket(id uuid.UUID) error
}

type service struct {
	repo     Repository
	rabbitmq *rabbit.RabbitMQ
}

func (s service) CreateTicket(ticketRequest CreateTicketRequest) (uuid.UUID, error) {
	var ticket Ticket

	ticket.ID = uuid.New()
	ticket.Email = ticketRequest.Email
	ticket.Title = ticketRequest.Title
	ticket.Body = ticketRequest.Body
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()

	uid, err := s.repo.Create(ticket)
	if err == nil {
		//publish rabbitmq message
	}
	return uid, err
}

func (s service) GetTicket(id uuid.UUID) (Ticket, error) {
	return s.repo.Get(id)
}

func (s service) DeleteTicket(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func NewService(r Repository, rabbit *rabbit.RabbitMQ) Service {
	return service{repo: r, rabbitmq: rabbit}
}
