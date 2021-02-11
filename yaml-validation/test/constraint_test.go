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

func TestPostConstraintValid(t *testing.T) {
	resp := httptest.NewRecorder()
	var constr = models.Constraint{
		Regex:    "[0-9]",
		Disabled: false,
	}

	b, _ := json.Marshal(constr)
	req, _ := http.NewRequest("POST", "/api/constraints/deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	router.ServeHTTP(resp, req)

	assert.Equal(t, 201, resp.Code)
}

func TestPostConstraintEnumAndRegexInvalid(t *testing.T) {
	resp := httptest.NewRecorder()
	var constr = models.Constraint{
		Regex:    "[0-9]",
		Enum:     []string{"1", "4"},
		Disabled: false,
	}

	b, _ := json.Marshal(constr)
	req, _ := http.NewRequest("POST", "/api/constraints/deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}

func TestPostConstraintEnumAndMinMaxInvalid(t *testing.T) {
	resp := httptest.NewRecorder()
	min := float32(1)
	max := float32(2)
	var constr = models.Constraint{
		Enum:     []string{"1", "4"},
		Min:      &min,
		Max:      &max,
		Disabled: false,
	}

	b, _ := json.Marshal(constr)
	req, _ := http.NewRequest("POST", "/api/constraints/deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}

func TestPostConstraintAllMissingInvalid(t *testing.T) {
	resp := httptest.NewRecorder()
	var constr = models.Constraint{
		Disabled: false,
	}

	b, _ := json.Marshal(constr)
	req, _ := http.NewRequest("POST", "/api/constraints/deployment-apps-v1/spec/replicas", bytes.NewBuffer(b))
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}

func TestPostConstraintInvalidPath(t *testing.T) {
	resp := httptest.NewRecorder()
	var constr = models.Constraint{
		Disabled: false,
		Regex:    "abc",
	}

	b, _ := json.Marshal(constr)
	req, _ := http.NewRequest("POST", "/api/constraints/abc-bcd/replicas", bytes.NewBuffer(b))
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}

func TestGetConstraintValid(t *testing.T) {
	resp := httptest.NewRecorder()
	var constr = models.Constraint{
		Enum:     []string{"1", "4"},
		Disabled: false,
	}
	b, _ := json.Marshal(constr)

	req, _ := http.NewRequest("POST", "/api/constraints/deployment-apps-v1/spec/minReadySeconds", bytes.NewBuffer(b))
	router.ServeHTTP(resp, req)

	req, _ = http.NewRequest("GET", "/api/constraints/deployment-apps-v1/spec", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var schema models.Schema
	json.Unmarshal(resp.Body.Bytes(), &schema)
	assert.Equal(t, 200, resp.Code)
	assert.NotNil(t, schema.Properties["minReadySeconds"].Constraint)

	if schema.Properties["minReadySeconds"].Constraint != nil {
		assert.Equal(t, len(schema.Properties["minReadySeconds"].Constraint.Enum), 2)
		assert.Equal(t, schema.Properties["minReadySeconds"].Constraint.Enum[0], "1")
	}
}

func TestGetConstraintInvalidPath(t *testing.T) {
	resp := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/constraints/abc/def", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}

func TestGetRootConstraints(t *testing.T) {
	resp := httptest.NewRecorder()
	var constr = models.Constraint{
		Enum:     []string{"1", "4"},
		Disabled: false,
	}
	b, _ := json.Marshal(constr)

	req, _ := http.NewRequest("POST", "/api/constraints/watchevent-v1", bytes.NewBuffer(b))
	router.ServeHTTP(resp, req)

	req, _ = http.NewRequest("GET", "/api/constraints", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var schemas []models.Schema
	json.Unmarshal(resp.Body.Bytes(), &schemas)
	var constraint *models.Constraint

	for _, schema := range schemas {
		if schema.GroupKindVersion[0].Kind == "WatchEvent" {
			constraint = schema.Constraint
		}
	}

	assert.Equal(t, 200, resp.Code)
	assert.NotNil(t, constraint)

	if constraint != nil {
		assert.Equal(t, len(constraint.Enum), 2)
		assert.Equal(t, constraint.Enum[0], "1")
	}
}

func TestGetConstraintOnIntInvalid(t *testing.T) {
	resp := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/constraints/deployment-apps-v1/spec/replicas", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}
