package services

import (
	"context"
	"yaml-validation/conf"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

// Sets up the database connection
// Connects to the database and sets mongoClient
func initDatabase() {
	credential := options.Credential{
		Username: conf.Configuration.Database.User,
		Password: conf.Configuration.Database.Password,
	}

	clientOpts := options.Client().
		ApplyURI(conf.Configuration.Database.Type + "://" + conf.Configuration.Database.Host + ":" + conf.Configuration.Database.Port).
		SetAuth(credential)

	mongoClient, _ = mongo.Connect(context.TODO(), clientOpts)
}

// Gets the mongoClient (connection to database)
// Returns: *mongo.Client: represent the database connection
func GetClient() *mongo.Client {
	return mongoClient
}
