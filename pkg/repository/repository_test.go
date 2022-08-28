package repository

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/karakaya/ticket/pkg/errors"
	"github.com/karakaya/ticket/pkg/model"
)

func TestInsertTicket(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		id := uuid.New()
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		mt.Coll = &mongo.Collection{}
		repo := NewRepository(mt.Client, mt.Coll)

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
		expectedUser := model.Ticket{
			ID:    uuid.New(),
			Title: "ticket title",
			Email: "email@ticket.com",
			Body:  "ticket@body",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "ticket.tickets", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedUser.ID},
			{Key: "title", Value: expectedUser.Title},
			{Key: "email", Value: expectedUser.Email},
			{Key: "body", Value: expectedUser.Body},
		}))

		repo := NewRepository(mt.Client, mt.Coll)

		response, err := repo.Get(expectedUser.ID)
		userResponse, _ := bson.Marshal(response)
		var ticket model.Ticket
		bson.Unmarshal(userResponse, &ticket)

		assert.Nil(t, err)
		assert.Equal(t, expectedUser, ticket)
	})
}

func TestDeleteTicket(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}})
		repo := NewRepository(mt.Client, mt.Coll)

		expected := errors.ErrNotFound

		err := repo.Delete(uuid.New())
		assert.Equal(t, err, expected)
	})
}
