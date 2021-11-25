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

func TestConstraintCreation_ConstraintWithRegexForStringProperty_Create(t *testing.T) {
	// Given
	regex := "[0-9]"
	var constraint = models.Constraint{
		Regex: &regex,
	}
	b, _ := json.Marshal(constraint)
	// When

	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/RuntimeClassList-node.k8s.io-v1beta1/metadata/continue", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)
	// Then
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintWithEnumForStringProperty_Create(t *testing.T) {
	// Given
	var constraint = models.Constraint{
		Enum: []string{"test", "case"},
	}
	b, _ := json.Marshal(constraint)

	// When

	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/template/metadata/clusterName", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
	request, _ = http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/spec/template/metadata/clusterName", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)
}

func TestConstraintCreation_ConstraintWithRegexAndEnumForStringProperty_Fail(t *testing.T) {
	// Given
	regex := "[0-9]"
	var constraint = models.Constraint{
		Regex: &regex,
		Enum:  []string{"1", "4"},
	}
	b, _ := json.Marshal(constraint)

	// When

	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/RuntimeClassList-node.k8s.io-v1beta1/metadata/continue", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintWithMinMaxForStringProperty_Fail(t *testing.T) {
	// Given
	min, max := float32(1), float32(2)
	var constraint = models.Constraint{
		Min: &min,
		Max: &max,
	}
	b, _ := json.Marshal(constraint)

	// When

	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/RuntimeClassList-node.k8s.io-v1beta1/metadata/continue", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintWithRegexAndMinMaxForStringProperty_Fail(t *testing.T) {
	// Given
	min, max := float32(1), float32(2)
	regex := "abc"
	var constraint = models.Constraint{
		Regex: &regex,
		Min:   &min,
		Max:   &max,
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/RuntimeClassList-node.k8s.io-v1beta1/metadata/continue", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintWithMinMaxForIntegerProperty_Create(t *testing.T) {
	// Given
	min, max := float32(1), float32(2)
	var constraint = models.Constraint{
		Min: &min,
		Max: &max,
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintWithEnumForIntegerProperty_Create(t *testing.T) {
	// Given
	var constraint = models.Constraint{
		Enum: []string{"1", "2"},
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintWithRegexForIntegerProperty_Create(t *testing.T) {
	// Given
	regex := "[0-9]"
	var constraint = models.Constraint{
		Regex: &regex,
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
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
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintWithRegexAndMinMaxForIntegerProperty_Fail(t *testing.T) {
	// Given
	min, max := float32(1), float32(2)
	regex := "abc"
	var constraint = models.Constraint{
		Regex: &regex,
		Min:   &min,
		Max:   &max,
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_EmptyConstraint_Fail(t *testing.T) {
	// Given
	var constraint = models.Constraint{}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_EmptyBody_Fail(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/replicas", nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_InvalidPath_Fail(t *testing.T) {
	// Given
	regex := "abc"
	var constraint = models.Constraint{
		Regex: &regex,
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/abc-bcd/replicas", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestConstraintCreation_ValidPathWithWrongCases_Fail(t *testing.T) {
	// Given
	min, max := float32(1), float32(2)
	var constraint = models.Constraint{
		Min: &min,
		Max: &max,
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintOnApiVersion_Fail(t *testing.T) {
	// Given
	regex := "abc"
	var constraint = models.Constraint{
		Regex: &regex,
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "http://localhost:8180/api/constraints/Deployment-apps-v1/apiVersion", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestConstraintCreation_ConstraintOnKind_Fail(t *testing.T) {
	// Given
	regex := "abc"
	var constraint = models.Constraint{
		Regex: &regex,
	}
	b, _ := json.Marshal(constraint)

	// When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "http://localhost:8180/api/constraints/Deployment-apps-v1/kind", bytes.NewBuffer(b))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", Token))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

// TODO: jaja
func TestConstraintCreation_IgnoreDisabledProperty_Create(t *testing.T) {
	// Given

	// When

	// Then

}
