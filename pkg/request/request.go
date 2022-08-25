package request

import (
	"time"

	"github.com/google/uuid"
)

type CreateTicketRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Email string `json:"email"`
}

type CreateTicketResponse struct {
	ID uuid.UUID `json:"_id"`
}

type UpdateTicketRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Email string `json:"email"`
}

type UpdateTicketResponse struct {
	ID        uuid.UUID `json:"_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindTicketResponse struct {
	ID        uuid.UUID `json:"_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
