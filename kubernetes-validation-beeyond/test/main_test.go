package test

import (
	"context"
	"github.com/Nerzal/gocloak/v10"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"kubernetes-validation-beeyond/conf"
	"kubernetes-validation-beeyond/routers"
	"os"
	"strings"
	"testing"
	"time"
)

var Router *gin.Engine
var Compose *testcontainers.LocalDockerCompose
var Token *jwt.Token

// Starts all creation and validation tests in a docker test-container
func TestMain(m *testing.M) {
	conf.Init()
	//_ = routers.Init()
	Router = routers.GetRouter()
	setupContainers()

	//Token, _ = middleware.ParseJwt(fetchTokenFromKeycloak(), middleware.FetchKeycloakPublicKey())

	code := m.Run()
	defer os.Exit(code)
	Compose.Down()
}

func setupContainers() {
	composeFilePaths := []string{"./resources/docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	Compose = testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)

	execError := Compose.
		WithCommand([]string{"up", "-d"}).
		Invoke()

	time.Sleep(10 * time.Second)

	services := Compose.Services
	idp := services["identity-provider"].(map[interface{}]interface{})
	conf.Configuration.Authentication.Port = strings.Split(idp["ports"].([]interface{})[0].(string), ":")[0]

	db := services["mongo-db"].(map[interface{}]interface{})
	conf.Configuration.Database.Port = strings.Split(db["ports"].([]interface{})[0].(string), ":")[0]

	err := execError.Error

	if err != nil {
		print(err)
	}

	//mongoDbContext := context.Background()
	//keycloakContext := context.Background()
	//req := testcontainers.ContainerRequest{
	//	Name:         "beeyond-mongo-db-test",
	//	Image:        "mongo:4.4.6",
	//	ExposedPorts: []string{conf.Configuration.Database.Port + "/tcp"},
	//	WaitingFor:   wait.ForLog("Waiting for connections"),
	//	Env: map[string]string{
	//		"MONGO_INITDB_DATABASE":      "beeyond_validation_db",
	//		"MONGO_INITDB_ROOT_USERNAME": "beeyond",
	//		"MONGO_INITDB_ROOT_PASSWORD": "beeyond"},
	//}
	//
	//mongoDbContainer, _ := testcontainers.GenericContainer(mongoDbContext, testcontainers.GenericContainerRequest{
	//	ContainerRequest: req,
	//	Started:          true,
	//})
	//
	//req = testcontainers.ContainerRequest{
	//	Name:         "beeyond-identity-provider-test",
	//	Image:        "jboss/keycloak:14.0.0",
	//	ExposedPorts: []string{conf.Configuration.Authentication.Port},
	//	Cmd:          []string{"-b2", "0.0.0.0", "-Dkeycloak.migration.action=import", "-Dkeycloak.profile.feature.upload_scripts=enabled", "-Dkeycloak.migration.provider=singleFile", "-Dkeycloak.migration.file=/tmp/school-realm.json"},
	//	//WaitingFor:   wait.ForHTTP("localhost:8180"),
	//	VolumeMounts: map[string]string{
	//		"./school-realm.json": "/tmp/school-realm.json",
	//		"./keycloak-theme":    "/opt/jboss/keycloak/themes/beeyond"},
	//	Env: map[string]string{
	//		"KEYCLOAK_USER2":    "beeyond",
	//		"KEYCLOAK_PASSWORD": "beeyond",
	//		"DB_VENDOR":         "H2"},
	//}
	//
	//keycloakContainer, _ := testcontainers.GenericContainer(keycloakContext, testcontainers.GenericContainerRequest{
	//	ContainerRequest: req,
	//	Started:          true,
	//})
	//
	//mongoDbContainer.Host(mongoDbContext)
	//host, err := keycloakContainer.Host(keycloakContext)
	//
	//print(host, err)
	//port, _ := mongoDbContainer.MappedPort(mongoDbContext, nat.Port(conf.Configuration.Database.Port))
	//conf.Configuration.Database.Port = strings.Split(string(port), "/")[0]
	//
	//port, _ = keycloakContainer.MappedPort(keycloakContext, nat.Port(conf.Configuration.Authentication.Port))
	//conf.Configuration.Authentication.Port = strings.Split(string(port), "/")[0]

}

func fetchTokenFromKeycloak() string {
	client := gocloak.NewClient(conf.Configuration.Authentication.Url + ":" + conf.Configuration.Authentication.Port)
	jwt, _ := client.Login(
		context.Background(),
		conf.Configuration.Authentication.ClientId,
		"",
		conf.Configuration.Authentication.Realm,
		conf.Configuration.Authentication.Username,
		conf.Configuration.Authentication.Password)

	return jwt.AccessToken
}
