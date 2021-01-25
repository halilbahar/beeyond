package test

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"yaml-validation/routers"

	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	router = routers.GetRouter()
}

func TestValidateEndpointShouldWork(t *testing.T) {
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}
