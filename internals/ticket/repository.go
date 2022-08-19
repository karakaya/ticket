package ticket

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ticket Ticket) (uuid.UUID, error)
	Get(id uuid.UUID) (Ticket, error)
	Delete(id uuid.UUID) error
}

type repository struct {
	client *mongo.Client
}

func (r repository) Get(id uuid.UUID) (Ticket, error) {
	collection := r.client.Database("ticket").Collection("tickets")
	var ticket Ticket
	collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&ticket)

	return ticket, nil
}

func (r repository) Create(ticket Ticket) (uuid.UUID, error) {
	collection := r.client.Database("ticket").Collection("tickets")
	_, err := collection.InsertOne(context.TODO(), ticket)
	if err != nil {
		return uuid.Nil, nil
	}
	return ticket.ID, nil
}
func (r repository) Delete(id uuid.UUID) error {
	collection := r.client.Database("ticket").Collection("tickets")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(mongoClient *mongo.Client) Repository {
	return repository{client: mongoClient}
}
