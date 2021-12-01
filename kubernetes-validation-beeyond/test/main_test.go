package test

import (
	"context"
	"github.com/Nerzal/gocloak/v10"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"kubernetes-validation-beeyond/conf"
	"kubernetes-validation-beeyond/middleware"
	"kubernetes-validation-beeyond/routers"
	"kubernetes-validation-beeyond/services"
	"os"
	"strings"
	"testing"
)

var Router *gin.Engine
var Compose *testcontainers.LocalDockerCompose
var Token *jwt.Token

// Starts all creation and validation tests in a docker test-container
func TestMain(m *testing.M) {
	conf.Init()
	setupContainers()
	//defer Compose.Down()
	services.Init()
	Router = routers.GetRouter()
	Token, _ = middleware.ParseJwt(fetchTokenFromKeycloak(), middleware.FetchKeycloakPublicKey())

	code := m.Run()
	os.Exit(code)
	Compose.Down()
}

func setupContainers() {
	composeFilePaths := []string{"./resources/docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	Compose = testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)

	execError := Compose.
		WithCommand([]string{"up", "-d"}).
		Invoke()

	//time.Sleep(10 * time.Second)

	services := Compose.Services
	idp := services["identity-provider"].(map[interface{}]interface{})
	conf.Configuration.Authentication.Port = strings.Split(idp["ports"].([]interface{})[0].(string), ":")[0]

	err := execError.Error

	if err != nil {
		print(err)
	}
}

func fetchTokenFromKeycloak() string {
	client := gocloak.NewClient(conf.Configuration.Authentication.Url + ":" + conf.Configuration.Authentication.Port)
	jwt, err := client.Login(
		context.Background(),
		conf.Configuration.Authentication.ClientId,
		"",
		conf.Configuration.Authentication.Realm,
		conf.Configuration.Authentication.Username,
		conf.Configuration.Authentication.Password)

	print(err)

	return jwt.AccessToken
}
