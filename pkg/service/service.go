package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/karakaya/ticket/pkg/model"
	"github.com/karakaya/ticket/pkg/rabbit"
	"github.com/karakaya/ticket/pkg/repository"
	"github.com/karakaya/ticket/pkg/request"
)

type Service interface {
	CreateTicket(ticketRequest request.CreateTicketRequest) (uuid.UUID, error)
	GetTicket(id uuid.UUID) (model.Ticket, error)
	DeleteTicket(id uuid.UUID) error
}

type service struct {
	repo repository.Repository
}

func (s service) CreateTicket(ticketRequest request.CreateTicketRequest) (uuid.UUID, error) {
	var ticket model.Ticket

	ticket.ID = uuid.New()
	ticket.Email = ticketRequest.Email
	ticket.Title = ticketRequest.Title
	ticket.Body = ticketRequest.Body
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()

	createResponse, err := s.repo.Create(ticket)
	if err == nil {
		ticketBytes, _ := ticket.Encoder()
		go rabbit.Publish(ticketBytes)
	}
	fmt.Println(createResponse)
	u, _ := createResponse.(model.Ticket)

	return u.ID, err
}

func (s service) GetTicket(id uuid.UUID) (model.Ticket, error) {
	ticket, err := s.repo.Get(id)
	if err != nil {
		return model.Ticket{}, err
	}
	ticketStruct, _ := ticket.(model.Ticket)
	return ticketStruct, nil

}

func (s service) DeleteTicket(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func NewService(r repository.Repository) Service {
	return service{repo: r}
}
