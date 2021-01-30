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

func TestPostConstraint(t *testing.T) {
	resp := httptest.NewRecorder()
	var constr = models.Constraint{
		Path:     "spec.replicas",
		Kind:     "Deployment",
		Regex:    "[0-9]",
		Disabled: false,
	}

	b, _ := json.Marshal(constr)
	req, _ := http.NewRequest("POST", "/api/constraints", bytes.NewBuffer(b))
	router.ServeHTTP(resp, req)

	assert.Equal(t, 201, resp.Code)
}

func TestGetConstraints(t *testing.T) {
	resp := httptest.NewRecorder()
	var constr = models.Constraint{
		Path:     "spec.selector.matchLabels.app",
		Kind:     "Deployment",
		Regex:    ".{1,7}",
		Disabled: false,
	}
	b, _ := json.Marshal(constr)

	req, _ := http.NewRequest("POST", "/api/constraints", bytes.NewBuffer(b))
	router.ServeHTTP(resp, req)

	req, _ = http.NewRequest("GET", "/api/constraints", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var constraints []models.Constraint
	json.Unmarshal(resp.Body.Bytes(), &constraints)
	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, constraints, constr)
	assert.Equal(t, len(constraints), 2)
}
