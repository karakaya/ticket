package ticket

import "github.com/google/uuid"

type Service interface {
	CreateTicket(ticketRequest CreateTicketRequest) (uuid.UUID, error)
	GetTicket(id uuid.UUID) (FindTicketResponse, error)
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

	return s.repo.Create(ticket)
}

func (s service) GetTicket(id uuid.UUID) (FindTicketResponse, error) {
	return FindTicketResponse{}, nil
}

func (s service) DeleteTicket(id uuid.UUID) error {
	return nil
}

func NewService(r Repository) Service {
	return service{repo: r}
}
