package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"kubernetes-validation-beeyond/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstraintToggle_DisableInteger_Valid(t *testing.T) {
	// Given
	min, max := float32(1), float32(2)
	var constraint = models.Constraint{
		Min:  &min,
		Max:  &max,
		Path: "spec.replicas",
		GroupKindVersion: models.GroupKindVersion{
			Group:   "apps",
			Kind:    "deployment",
			Version: "v1",
		},
	}
	b, _ := json.Marshal(constraint)

	// When
	err := models.SaveConstraint(constraint)

	// Then
	assert.Nil(t, err)

	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Given When
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/constraints/Deployment-apps-v1/spec", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.True(t, schema.Properties["replicas"].Constraint.Disabled)
}

func TestConstraintToggle_EnableInteger_Valid(t *testing.T) {
	// Given
	min, max := float32(1), float32(2)
	var constraint = models.Constraint{
		Min:      &min,
		Max:      &max,
		Path:     "spec.replicas",
		Disabled: true,
		GroupKindVersion: models.GroupKindVersion{
			Group:   "apps",
			Kind:    "deployment",
			Version: "v1",
		},
	}
	b, _ := json.Marshal(constraint)

	// When
	err := models.SaveConstraint(constraint)

	// Then
	assert.Nil(t, err)

	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Given When
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/constraints/Deployment-apps-v1/spec", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.False(t, schema.Properties["replicas"].Constraint.Disabled)
}

func TestConstraintToggle_DisableString_Valid(t *testing.T) {
	// Given
	enum := []string{"ja", "nein"}
	var constraint = models.Constraint{
		Enum:     enum,
		Path:     "metadata.clusterName",
		Disabled: false,
		GroupKindVersion: models.GroupKindVersion{
			Group:   "apps",
			Kind:    "deployment",
			Version: "v1",
		},
	}
	b, _ := json.Marshal(constraint)

	// When
	err := models.SaveConstraint(constraint)

	// Then
	assert.Nil(t, err)

	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/metadata/clusterName", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Given When
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/constraints/Deployment-apps-v1/metadata", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.True(t, schema.Properties["clusterName"].Constraint.Disabled)
}

func TestConstraintToggle_EnableString_Valid(t *testing.T) {
	// Given
	enum := []string{"ja", "nein"}
	var constraint = models.Constraint{
		Enum:     enum,
		Path:     "metadata.clusterName",
		Disabled: true,
		GroupKindVersion: models.GroupKindVersion{
			Group:   "apps",
			Kind:    "deployment",
			Version: "v1",
		},
	}
	b, _ := json.Marshal(constraint)

	// When
	err := models.SaveConstraint(constraint)

	// Then
	assert.Nil(t, err)

	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/metadata/clusterName", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Given When
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/constraints/Deployment-apps-v1/metadata", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.False(t, schema.Properties["clusterName"].Constraint.Disabled)
}

func TestConstraintToggle_DisableObject_Valid(t *testing.T) {
	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/spec", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Given
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/constraints/Deployment-apps-v1", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	//assert.True(t, schema.Properties["spec"].Constraint.Disabled)
}

func TestConstraintToggle_EnableObject_Valid(t *testing.T) {
	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/metadata", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Given
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/metadata", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Given
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/constraints/Deployment-apps-v1", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	fmt.Println(schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.False(t, schema.Properties["metadata"].Constraint.Disabled)
}

func TestConstraintToggle_DisableRootObject_Valid(t *testing.T) {
	// Given
	enum := []string{"ja", "nein"}
	var constraint = models.Constraint{
		Enum:     enum,
		Path:     "",
		Disabled: false,
		GroupKindVersion: models.GroupKindVersion{
			Group:   "apps",
			Kind:    "deployment",
			Version: "v1",
		},
	}
	b, _ := json.Marshal(constraint)

	// When
	err := models.SaveConstraint(constraint)

	// Then
	assert.Nil(t, err)

	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Given When
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/constraints/Deployment-apps-v1", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	fmt.Println(schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.True(t, schema.Constraint.Disabled)
}

func TestConstraintToggle_EnableRootObject_Valid(t *testing.T) {
	// Given
	enum := []string{"ja", "nein"}
	var constraint = models.Constraint{
		Enum:     enum,
		Path:     "",
		Disabled: true,
		GroupKindVersion: models.GroupKindVersion{
			Group:   "apps",
			Kind:    "Deployment",
			Version: "v1",
		},
	}
	b, _ := json.Marshal(constraint)

	// When
	err := models.SaveConstraint(constraint)

	// Then
	assert.Nil(t, err)

	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Given When
	responseRecorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/constraints/Deployment-apps-v1", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	var schema models.Schema
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &schema)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.False(t, schema.Constraint.Disabled)
}

func TestConstraintToggle_DisableRequiredString_Fail(t *testing.T) {
	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/MutatingWebhookConfiguration-admissionregistration.k8s.io-v1/webhooks/name", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintToggle_DisableRequiredInteger_Fail(t *testing.T) {
	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/ControllerRevision-apps-v1beta1/revision", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintToggle_DisableRequiredObject_Fail(t *testing.T) {
	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/APIGroup-v1/name", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintToggle_DisableApiVersion_Fail(t *testing.T) {
	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/apiVersion", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestConstraintToggle_DisableKind_Fail(t *testing.T) {
	// Given
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/kind", nil)
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}
