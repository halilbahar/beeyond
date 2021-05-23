package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"kubernetes-validation-beeyond/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstraintGet_WithConstraint_Valid(t *testing.T) {
	// Given
	var constraint = models.Constraint{
		Enum: []string{"1", "4"},
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/minReadySeconds", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)

	// Given When
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/constraints/Deployment-apps-v1/spec", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotNil(t, schema.Properties["minReadySeconds"].Constraint)
	assert.Equal(t, len(schema.Properties["minReadySeconds"].Constraint.Enum), 2)
	assert.Equal(t, schema.Properties["minReadySeconds"].Constraint.Enum[0], "1")
	assert.Equal(t, schema.Properties["minReadySeconds"].Constraint.Enum[1], "4")
}

func TestConstraintGet_WithoutConstraint_Valid(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/constraints/Deployment-apps-v1/spec", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestConstraintGet_ShouldNotContainKind(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/constraints/Deployment-apps-v1/spec", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Nil(t, schema.Properties["kind"])
}

func TestConstraintGet_ShouldNotContainApiVersion(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/constraints/Deployment-apps-v1/spec", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Nil(t, schema.Properties["apiVersion"])
}

func TestConstraintGet_RootElementShouldNotContainKind(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/constraints", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schemas []*models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schemas)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Nil(t, schemas[0].Properties["kind"])
}

func TestConstraintGet_RootElementShouldNotContainApiVersion(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/constraints", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schemas []*models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schemas)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Nil(t, schemas[0].Properties["apiVersion"])
}

func TestConstraintGet_InvalidPath_Fail(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/constraints/abc/def", nil)
	responseRecorder = httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestConstraintGet_IntegerPropertyPath_Fail(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/constraints/deployment-apps-v1/spec/replicas", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestConstraintGet_StringPropertyPath_Fail(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/constraints/RuntimeClassList-node.k8s.io-v1beta1/metadata/continue", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestConstraintGet_IntOrStringPath_Fail(t *testing.T) {
	// io.k8s.apimachinery.pkg.util.intstr.IntOrString
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/constraints/PodDisruptionBudget-policy-v1beta1/spec/maxUnavailable", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestConstraintGet_RootElement_Valid(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/constraints/Deployment-apps-v1", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestConstraintGet_RootElementWithConstraint_Valid(t *testing.T) {
	// Given
	var constr = models.Constraint{
		Enum: []string{"1", "4"},
	}
	b, _ := json.Marshal(constr)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/WatchEvent-v1", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)

	// Given
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/constraints", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	var schemas []models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schemas)
	var constraint *models.Constraint

	for _, schema := range schemas {
		if schema.GroupKindVersion[0].Kind == "WatchEvent" {
			constraint = schema.Constraint
		}
	}

	assert.NotNil(t, constraint)
	assert.Equal(t, len(constraint.Enum), 2)
	assert.Equal(t, constraint.Enum[0], "1")
	assert.Equal(t, constraint.Enum[1], "4")
}

func TestConstraintGet_RootElementWithoutApiVersionAndKind_Valid(t *testing.T) {

}
