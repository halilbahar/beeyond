package test

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
	"yaml-validation/pkg/setting"
	"yaml-validation/routers"
	"yaml-validation/services"
)

var Router *gin.Engine
var mongoDbContainer testcontainers.Container
var mongoDbContext context.Context

func TestMain(m *testing.M) {
	setting.Init()
	services.Init()

	//setupMongoDbContainer()
	Router = routers.GetRouter()

	code := m.Run()
	//mongoDbContainer.Terminate(mongoDbContext)
	os.Exit(code)
}

func setupMongoDbContainer() {
	mongoDbContext = context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongodb",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections"),
	}

	mongoDbContainer, _ = testcontainers.GenericContainer(mongoDbContext, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	//_, _ = mongoDbContainer.Host(mongoDbContext)
	//_, _ = mongoDbContainer.MappedPort(mongoDbContext, "27017")
}
