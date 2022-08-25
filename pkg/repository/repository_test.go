package repository

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/karakaya/ticket/pkg/model"
)

func TestInsertTicket(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		// ticketCollection := mt.Coll
		id := uuid.New()
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		repo := NewRepository(mt.Client)

		_, err := repo.Create(model.Ticket{
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

		expectedTicket := model.Ticket{
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

func TestDeleteTicket(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		repo := NewRepository(mt.Client)

		err := repo.Delete(uuid.New())
		assert.Nil(t, err)
	})
}
