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
	GetAll() (interface{}, error)
	Delete(id uuid.UUID) error
}

type repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (r repository) Create(ticket interface{}) (interface{}, error) {

	_, err := r.collection.InsertOne(context.TODO(), ticket)
	if err != nil {
		return uuid.Nil, nil
	}
	return ticket, nil
}

func (r repository) Get(id uuid.UUID) (interface{}, error) {

	var ticket interface{}
	result := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&ticket)
	if result == mongo.ErrNoDocuments {
		return nil, er.ErrNotFound
	}
	return ticket, nil
}

func (r repository) GetAll() (interface{}, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	var dataArr bson.A
	for cursor.Next(context.TODO()) {
		var data bson.M
		cursor.Decode(&data)
		dataArr = append(dataArr, data)
	}

	return dataArr, nil
}

func (r repository) Delete(id uuid.UUID) error {

	result := r.collection.FindOneAndDelete(context.TODO(), bson.M{"_id": id})

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
