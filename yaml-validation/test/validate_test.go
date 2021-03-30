package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"yaml-validation/models"
)

func TestValidateEndpointShouldWork(t *testing.T) {
	models.DeleteAll()
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("test/resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}

func TestValidateEndpointShouldReturnError(t *testing.T) {
	models.DeleteAll()
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("test/resources/invalid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	assert.Equal(t, 422, resp.Code)
}

func TestValidateEndpointWithValidConstraint_ShouldWork(t *testing.T) {
	// given
	gkv := models.GroupKindVersion{
		Group:   "apps",
		Kind:    "Deployment",
		Version: "v1",
	}
	models.DeleteAll()
	enum := []string{"1", "3"}
	var constraint = models.Constraint{
		Enum:             enum,
		Path:             "spec.replicas",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// when
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("test/resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// then
	assert.Equal(t, 200, resp.Code)
	models.DeleteConstraint("spec.replicas", gkv)
}

func TestValidateEndpointWithInvalidConstraint_ShouldReturnError(t *testing.T) {
	// given
	gkv := models.GroupKindVersion{
		Group:   "apps",
		Kind:    "Deployment",
		Version: "v1",
	}
	models.DeleteAll()
	enum := []string{"1", "2"}
	var constraint = models.Constraint{
		Enum:             enum,
		Path:             "spec.replicas",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// when
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("test/resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// then
	assert.Equal(t, 422, resp.Code)
	models.DeleteConstraint("spec.replicas", gkv)
}
