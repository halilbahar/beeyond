package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"yaml-validation/models"
)

func TestConstraintCreation_ConstraintWithRegexForStringProperty_Create(t *testing.T) {
	// Given
	var constraint = models.Constraint{
		Regex: "[0-9]",
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	// TODO: change replicas (integer) to something with string
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintWithEnumForStringProperty_Create(t *testing.T) {
	// Given
	// When
	// Then
}

func TestConstraintCreation_ConstraintWithRegexAndEnumForStringProperty_Fail(t *testing.T) {
	// Given
	var constraint = models.Constraint{
		Regex: "[0-9]",
		Enum:  []string{"1", "4"},
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	// TODO: change replicas (integer) to something with string
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintWithMinMaxForStringProperty_Fail(t *testing.T) {
	// Strings can only have enum and regex
	// Given
	// When
	// Then
}

func TestConstraintCreation_ConstraintWithRegexAndMinMaxForStringProperty_Fail(t *testing.T) {
	// Given
	// When
	// Then
}

func TestConstraintCreation_ConstraintWithMinMaxForIntegerProperty_Create(t *testing.T) {
	// Given
	// When
	// Then
}

func TestConstraintCreation_ConstraintWithEnumForIntegerProperty_Create(t *testing.T) {
	// Given
	// When
	// Then
}

func TestConstraintCreation_ConstraintWithEnumAndMinMax_Fail(t *testing.T) {
	// Given
	min, max := float32(1), float32(2)
	var constraint = models.Constraint{
		Enum: []string{"1", "4"},
		Min:  &min,
		Max:  &max,
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintWithRegexAndMinMaxForIntegerProperty_Fail(t *testing.T) {
	// Given
	// When
	// Then
}

func TestConstraintCreation_EmptyConstraint_Fail(t *testing.T) {
	// Given
	var constraint = models.Constraint{}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-Apps-v1/spec/replicas", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_EmptyBody_Fail(t *testing.T) {
	// Given
	// When
	// Then
}

func TestConstraintCreation_InvalidPath_Fail(t *testing.T) {
	// Given
	var constraint = models.Constraint{
		Regex: "abc",
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/abc-bcd/replicas", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestConstraintCreation_ValidPathWithWrongCases_Fail(t *testing.T) {
	// Given
	// When
	// Then
}

func TestConstraintCreation_ConstraintWith(t *testing.T) {
	// Given
	// When
	// Then
}

func TestConstraintGet_Valid(t *testing.T) {
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
	// Given
	// When
	// Then
}

func TestConstraintGet_IntOrStringPath_Fail(t *testing.T) {
	// io.k8s.apimachinery.pkg.util.intstr.IntOrString
	// Given
	// When
	// Then
}

func TestConstraintGet_RootElement_Valid(t *testing.T) {
	// Given
	// When
	// Then
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
