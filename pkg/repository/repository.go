package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/karakaya/ticket/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ticket interface{}) (interface{}, error)
	Get(id uuid.UUID) (interface{}, error)
	Delete(id uuid.UUID) error
}

type repository struct {
	client *mongo.Client
}

func (r repository) Get(id uuid.UUID) (interface{}, error) {
	collection := r.client.Database("ticket").Collection("tickets")
	var ticket model.Ticket
	collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&ticket)

	return ticket, nil
}

func (r repository) Create(ticket interface{}) (interface{}, error) {
	collection := r.client.Database("ticket").Collection("tickets")
	_, err := collection.InsertOne(context.TODO(), ticket)
	if err != nil {
		return uuid.Nil, nil
	}
	return ticket, nil
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
