package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"yaml-validation/pkg/setting"
	"yaml-validation/routers"
	"yaml-validation/services"

	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
var compose *testcontainers.LocalDockerCompose

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	compose.Down()
	os.Exit(code)
}

func setup() {
	setting.Init()
	services.Init()

	composeFilePaths := []string{"conf/docker-compose.yaml"}
	identifier := strings.ToLower(uuid.New().String())

	compose = testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	execError := compose.
		WithCommand([]string{"up", "-d"}).
		WithEnv(map[string]string{
			"key1": "value1",
			"key2": "value2",
		}).
		Invoke()
	err := execError.Error
	if err != nil {
		fmt.Errorf("Could not run compose file: %v - %v", composeFilePaths, err)
	}

	router = routers.GetRouter()
}

func TestValidateEndpointShouldWork(t *testing.T) {
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}

func TestValidateEndpointShouldReturnError(t *testing.T) {
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("resources/invalid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	router.ServeHTTP(resp, req)

	assert.Equal(t, 422, resp.Code)
}
