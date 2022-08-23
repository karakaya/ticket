package ticket

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsertTicket(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
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

func TestFindTicket(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {

		expectedTicket := Ticket{
			ID:    uuid.New(),
			Title: "ticket title",
			Email: "john.doe@test.com",
			Body:  "body title",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "ticket.tickets", mtest.FirstBatch, bson.D{
			{"_id", expectedTicket.ID},
			{"title", expectedTicket.Title},
			{"email", expectedTicket.Email},
			{"body", expectedTicket.Body},
		}))

		repo := NewRepository(mt.Client)

		ticketResponse, err := repo.Get(expectedTicket.ID)
		assert.Nil(t, err)

		assert.Equal(t, expectedTicket, ticketResponse)
	})
}
