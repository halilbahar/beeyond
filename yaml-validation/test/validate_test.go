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

func TestValidateEndpoint_ShouldWork(t *testing.T) {
	// Given
	models.DeleteAll()

	// When
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("test/resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 200, resp.Code)
}

func TestValidateEndpoint_ShouldReturnError(t *testing.T) {
	// Given
	models.DeleteAll()

	// When
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("test/resources/invalid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 422, resp.Code)
}

func TestValidateEndpoint_WithValidConstraint_ShouldWork(t *testing.T) {
	// Given
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

	// When
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("test/resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 200, resp.Code)
	models.DeleteConstraint("spec.replicas", gkv)
}

func TestValidateEndpoint_WithInvalidConstraint_ShouldReturnError(t *testing.T) {
	// Given
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

	// When
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("test/resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 422, resp.Code)
	models.DeleteConstraint("spec.replicas", gkv)
}

/////////////////////////////
// Min & Max single values //
/////////////////////////////

func TestValidateEndpoint_WithMinMaxAndIntegerValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithMinMaxAndFloatValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithMinMaxAndStringValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithMinMaxAndBooleanValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithMinMaxAndObjectValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

////////////////////////////
// Min & Max array values //
////////////////////////////

func TestValidateEndpoint_WithMinMaxAndIntegerArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithMinMaxAndFloatArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithMinMaxAndStringArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithMinMaxAndBooleanArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithMinMaxAndObjectArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

////////////////////////
// Enum single values //
////////////////////////

func TestValidateEndpoint_WithEnumAndIntegerValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithEnumAndFloatValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithEnumAndStringValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithEnumAndBooleanValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithEnumAndObjectValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

///////////////////////
// Enum array values //
///////////////////////

func TestValidateEndpoint_WithEnumAndIntegerArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithEnumAndFloatArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithEnumAndStringArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithEnumAndBooleanArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithEnumAndObjectArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

/////////////////////////
// Regex single values //
/////////////////////////

func TestValidateEndpoint_WithRegexAndIntegerValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndFloatValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndStringValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndBooleanValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndObjectValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

////////////////////////
// Regex array values //
////////////////////////

func TestValidateEndpoint_WithRegexAndIntegerArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndFloatArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndStringArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndBooleanArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndObjectArrayValue_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}
