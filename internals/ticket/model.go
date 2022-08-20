package ticket

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID    uuid.UUID `bson:"_id" json:"id"`
	Title string    `bson:"title" json:"title"`
	Body  string    `bson:"body" json:"body"`
	// User      User
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type User struct {
	ID        uuid.UUID `bson:"_id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
}

func (t *Ticket) Encoder() ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(t)
	return b.Bytes(), err
}
