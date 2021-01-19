package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func initDatabase() {
	credential := options.Credential{
		Username: "beeyond",
		Password: "beeyond",
	}

	clientOpts := options.Client().
		ApplyURI("mongodb://localhost:27017").
		SetAuth(credential)

	mongoClient, _ = mongo.Connect(context.TODO(), clientOpts)
}

func GetClient() *mongo.Client {
	return mongoClient
}
