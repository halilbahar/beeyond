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
	mongoDbContext := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongo",
		ExposedPorts: []string{conf.Configuration.Database.Port+"/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections"),
		Env: map[string]string{
			"MONGO_INITDB_DATABASE":"beeyond_validation_db",
			"MONGO_INITDB_ROOT_USERNAME":"beeyond",
			"MONGO_INITDB_ROOT_PASSWORD":"beeyond"},
	}
	fmt.Println("---------------------------------------------CONF VALUE: "+conf.Configuration.Database.Port)

	mongoDbContainer, _ = testcontainers.GenericContainer(mongoDbContext, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	_, _ = mongoDbContainer.Host(mongoDbContext)
	port, _ := mongoDbContainer.MappedPort(mongoDbContext, nat.Port(conf.Configuration.Database.Port))

	conf.Configuration.Database.Port = strings.Split(string(port), "/")[0]

	//conf.V.SetEnvPrefix(conf.EnvVarPrefix)
	//conf.V.BindEnv(conf.EnvVarBindVar)
	//
	//os.Setenv(conf.EnvVar, strings.Split(string(port), "/")[0])
	//v := os.Getenv(conf.EnvVar)
	//fmt.Println("-------------------------------------ENV VAR VALUE: "+v)
}

