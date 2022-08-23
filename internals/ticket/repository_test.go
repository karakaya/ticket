package ticket

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsertTicket(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("insert ticket success", func(mt *mtest.T) {
		// ticketCollection := mt.Coll
		id := uuid.New()
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		repo := NewRepository(mt.Client)

		_, err := repo.Create(Ticket{
			ID:        id,
			Title:     "ticket title",
			Body:      "body title",
			Email:     "email@sample.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		assert.Nil(t, err)
	})
}
