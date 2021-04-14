package test

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"strings"
	"testing"
	"yaml-validation/conf"
	"yaml-validation/routers"
	"yaml-validation/services"
)

var Router *gin.Engine
var mongoDbContainer testcontainers.Container

func TestMain(m *testing.M) {
	conf.Init()

	setupMongoDbContainer()

	services.Init()

	Router = routers.GetRouter()

	code := m.Run()
	mongoDbContainer.Terminate(context.Background())
	os.Exit(code)
}

func setupMongoDbContainer() {
	//credential := options.Credential{
	//	Username: conf.Configuration.Database.User,
	//	Password: conf.Configuration.Database.Password,
	//}
	//
	//clientOpts := options.Client().
	//	ApplyURI(conf.Configuration.Database.Type + "://" + conf.Configuration.Database.Host + ":" + conf.Configuration.Database.Port).
	//	SetAuth(credential)
	//
	//client, err := mongo.Connect(context.TODO(), clientOpts)
	//
	//if err == nil{
	//	client.Disconnect(context.TODO())
	//	return
	//}

	mongoDbContext := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongo",
		ExposedPorts: []string{conf.Configuration.Database.Port + "/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections"),
		Env: map[string]string{
			"MONGO_INITDB_DATABASE":      "beeyond_validation_db",
			"MONGO_INITDB_ROOT_USERNAME": "beeyond",
			"MONGO_INITDB_ROOT_PASSWORD": "beeyond"},
	}

	mongoDbContainer, _ = testcontainers.GenericContainer(mongoDbContext, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	_, _ = mongoDbContainer.Host(mongoDbContext)
	port, _ := mongoDbContainer.MappedPort(mongoDbContext, nat.Port(conf.Configuration.Database.Port))

	conf.Configuration.Database.Port = strings.Split(string(port), "/")[0]

	fmt.Println("------------------------------------------test db port: " + conf.Configuration.Database.Port)
}
