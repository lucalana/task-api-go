package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm/utils"
)

type MongoRepositoryContext struct {
	*mongo.Collection
	client *mongo.Client
}

func NewMongoRepositoryContext(uri, dbName, collectionName string) (*MongoRepositoryContext, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, utils.InternalServerError(fmt.Sprintf("erro ao conectar com mongoDB: %v", err))
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, utils.InternalServerError(fmt.Sprintf("erro ao testar conexão com mongoDB: %v", err))
	}
	collection := client.Database(dbName).Collection(collectionName)

	return &MongoRepositoryContext{
		Collection: collection,
		client:     client,
	}, nil
}

func (r *MongoRepositoryContext) Create(contextServer context.Context, document interface{}) error {
	_, err := r.Collection.InsertOne(contextServer, document)
	if err != nil {
		return utils.BadRequestError(fmt.Sprintf("erro ao inserir documento: %v", err))
	}
	return nil
}
