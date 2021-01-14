package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Constraint struct {
	Path string `json:"path"`
	Kind string `json:"kind"`
	Regex string `json:"regex"`
	Disabled bool `json:"disabled"`
}

var mongoClient *mongo.Client

func connect() {
	var err error
	credential := options.Credential{
		Username: "beeyond",
		Password: "beeyond",
	}

	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	mongoClient, err = mongo.Connect(context.TODO(), clientOpts)

	PrintError(err)
}

func disconnect() {
	PrintError(mongoClient.Disconnect(context.TODO()))
}

func SaveConstraint(constraint Constraint) error {
	connect()
	collection := mongoClient.Database("beeyond_db").Collection("Constraints")
	_, err := collection.InsertOne(context.TODO(), constraint)
	fmt.Print(err)
	disconnect()
	return err
}

func PrintError(err error) {
	if err != nil{
		panic(err)
	}
}