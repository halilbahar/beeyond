package services

import (
	"context"

	"yaml-validation/pkg/setting"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func initDatabase() {
	credential := options.Credential{
		Username: setting.DatabaseSetting.User,
		Password: setting.DatabaseSetting.Password,
	}

	clientOpts := options.Client().
		ApplyURI(setting.DatabaseSetting.Type + "://" + setting.DatabaseSetting.Host).
		SetAuth(credential)

	mongoClient, _ = mongo.Connect(context.TODO(), clientOpts)
}

func GetClient() *mongo.Client {
	return mongoClient
}
