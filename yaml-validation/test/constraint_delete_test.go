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

func TestConstraintDelete_ValidConstraint_Valid(t *testing.T) {
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
	request, _ := http.NewRequest("DELETE", "/api/constraints/Deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNoContent, responseRecorder.Code)
}

func TestConstraintDelete_NoConstraint_Fail(t *testing.T) {
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
	request, _ := http.NewRequest("DELETE", "/api/constraints/Deployment-apps-v1/spec/paused", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintDelete_InvalidPath_Fail(t *testing.T) {
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
	request, _ := http.NewRequest("DELETE", "/api/constraints/Deployment-apps-v1/spec/replica", bytes.NewBuffer(b))
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}
