package repository

import (
	"context"

	"github.com/google/uuid"
	er "github.com/karakaya/ticket/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ticket interface{}) (interface{}, error)
	Get(id uuid.UUID) (interface{}, error)
	GetAll() ([]interface{}, error)
	Delete(id uuid.UUID) error
}

type repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (r repository) Create(ticket interface{}) (interface{}, error) {
	collection := r.client.Database("ticket").Collection("tickets")
	_, err := collection.InsertOne(context.TODO(), ticket)
	if err != nil {
		return uuid.Nil, nil
	}
	return ticket, nil
}

func (r repository) Get(id uuid.UUID) (interface{}, error) {
	collection := r.client.Database("ticket").Collection("tickets")
	var ticket interface{}
	result := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&ticket)
	if result == mongo.ErrNoDocuments {
		return nil, er.ErrNotFound
	}
	return ticket, nil
}

func (r repository) GetAll() ([]interface{}, error) {
	result := make([]interface{}, 0)
	allData, _ := r.collection.Find(context.TODO(), bson.D{})
	for allData.Next(context.TODO()) {
		var element interface{}
		allData.Decode(&element)
		result = append(result, element)
	}
	return result, nil

}

func (r repository) Delete(id uuid.UUID) error {
	collection := r.client.Database("ticket").Collection("tickets")
	result := collection.FindOneAndDelete(context.TODO(), bson.M{"_id": id})

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return er.ErrNotFound
		}
		return result.Err()
	}
	return nil
}

func NewRepository(mongoClient *mongo.Client, collection *mongo.Collection) Repository {
	return repository{client: mongoClient, collection: collection}
}
