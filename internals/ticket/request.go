package ticket

import "github.com/google/uuid"

type CreateTicketRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Email string `json:"email"`
}

type CreateTicketResponse struct {
	ID uuid.UUID `json:"id"`
}

type FindTicketResponse struct {
}
