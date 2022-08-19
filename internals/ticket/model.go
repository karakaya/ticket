package ticket

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID    uuid.UUID `bson:"_id"`
	Title string    `bson:"title"`
	Body  string    `bson:"body"`
	// User      User
	Email     string `bson:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        uuid.UUID `bson:"_id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time
}
