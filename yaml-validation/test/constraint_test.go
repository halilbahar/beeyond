package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostConstraint(t *testing.T) {
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/constraints", bytes.NewBufferString(`{"kind": "Deployment", "path": "spec.replicas", "regex": "[0-9]", "disabled": false}`))
	router.ServeHTTP(resp, req)

	assert.Equal(t, 201, resp.Code)
}
