package ticket

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID        uuid.UUID
	Title     string
	Body      string
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
