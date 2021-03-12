package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidateEndpointShouldWork(t *testing.T) {
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("test/resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}

func TestValidateEndpointShouldReturnError(t *testing.T) {
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("test/resources/invalid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	assert.Equal(t, 422, resp.Code)
}
