package services

import (
	"../pkg/setting"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func Init() {
	credential := options.Credential{
		Username: setting.DatabaseSetting.User,
		Password: setting.DatabaseSetting.Password,
	}

	// setting.DatabaseSetting.Type + "://" + setting.DatabaseSetting.Host
	clientOpts := options.Client().ApplyURI(setting.DatabaseSetting.Type + "://" + setting.DatabaseSetting.Host).SetAuth(credential)
	mongoClient, _ = mongo.Connect(context.TODO(), clientOpts)
}

func GetClient() *mongo.Client {
	return mongoClient
}
