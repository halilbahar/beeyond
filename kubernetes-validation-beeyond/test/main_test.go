package test

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"kubernetes-validation-beeyond/conf"
	"kubernetes-validation-beeyond/routers"
	"kubernetes-validation-beeyond/services"
	"os"
	"strings"
	"testing"
)

var Router *gin.Engine
var mongoDbContainer testcontainers.Container
var Token *jwt.Token

// Starts all creation and validation tests in a docker test-container
func TestMain(m *testing.M) {
	conf.Init()
	setupMongoDbContainer()
	services.Init()
	Router = routers.GetRouter()

	code := m.Run()
	mongoDbContainer.Terminate(context.Background())
	os.Exit(code)
}

// Sets up a testcontainer (docker)
func setupMongoDbContainer() {
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
}
