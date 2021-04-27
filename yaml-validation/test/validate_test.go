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
	c, _ := ioutil.ReadFile("./resources/valid.yaml")
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
	c, _ := ioutil.ReadFile("./resources/invalid.yaml")
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
	c, _ := ioutil.ReadFile("./resources/valid.yaml")
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
	c, _ := ioutil.ReadFile("./resources/valid.yaml")
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
	models.DeleteAll()
	min, max := float32(1), float32(4)
	gkv := models.GroupKindVersion{
		Group:   "apps",
		Kind:    "Deployment",
		Version: "v1",
	}
	var constraint = models.Constraint{
		Min:              &min,
		Max:              &max,
		Path:             "spec.replicas",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When

	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("./resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 200, resp.Code)
	models.DeleteConstraint("spec.replicas", gkv)

	// Given
	max = float32(2)
	constraint = models.Constraint{
		Min:              &min,
		Max:              &max,
		Path:             "spec.replicas",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 422, resp.Code)
	models.DeleteConstraint("spec.replicas", gkv)
}

////////////////////////////
// Min & Max array values //
////////////////////////////

func TestValidateEndpoint_WithMinMaxAndIntegerArrayValue_ShouldWork(t *testing.T) {
	// Given
	models.DeleteAll()
	min, max := float32(1), float32(4)
	gkv := models.GroupKindVersion{
		Group:   "",
		Kind:    "Pod",
		Version: "v1",
	}
	var constraint = models.Constraint{
		Min:              &min,
		Max:              &max,
		Path:             "spec.securityContext.supplementalGroups",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When

	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("./resources/validIntegerArray.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 200, resp.Code)
	models.DeleteConstraint("spec.securityContext.supplementalGroups", gkv)

	// Given
	max = float32(2)
	constraint = models.Constraint{
		Min:              &min,
		Max:              &max,
		Path:             "spec.securityContext.supplementalGroups",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 422, resp.Code)
	models.DeleteConstraint("spec.securityContext.supplementalGroups", gkv)
}

////////////////////////
// Enum single values //
////////////////////////

func TestValidateEndpoint_WithEnumAndIntegerValue_ShouldWork(t *testing.T) {
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
	c, _ := ioutil.ReadFile("./resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 200, resp.Code)
	models.DeleteConstraint("spec.replicas", gkv)

	// Given
	enum = []string{"1", "2"}
	constraint = models.Constraint{
		Enum:             enum,
		Path:             "spec.replicas",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 422, resp.Code)
	models.DeleteConstraint("spec.replicas", gkv)
}

func TestValidateEndpoint_WithEnumAndStringValue_ShouldWork(t *testing.T) {
	// Given
	gkv := models.GroupKindVersion{
		Group:   "apps",
		Kind:    "Deployment",
		Version: "v1",
	}
	models.DeleteAll()
	enum := []string{"beeyond", "isAwesome"}
	var constraint = models.Constraint{
		Enum:             enum,
		Path:             "metadata.clusterName",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("./resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 200, resp.Code)
	models.DeleteConstraint("metadata.clusterName", gkv)

	// Given
	enum = []string{"beyond", "isAwesome"}
	constraint = models.Constraint{
		Enum:             enum,
		Path:             "metadata.clusterName",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 422, resp.Code)
	models.DeleteConstraint("metadata.clusterName", gkv)
}

func TestValidateEndpoint_WithEnumAndBooleanValue_ShouldWork(t *testing.T) {
	gkv := models.GroupKindVersion{
		Group:   "apps",
		Kind:    "Deployment",
		Version: "v1",
	}
	models.DeleteAll()
	enum := []string{"false"}
	var constraint = models.Constraint{
		Enum:             enum,
		Path:             "spec.paused",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("./resources/valid.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 200, resp.Code)
	models.DeleteConstraint("spec.paused", gkv)

	// Given
	enum = []string{"true"}
	constraint = models.Constraint{
		Enum:             enum,
		Path:             "spec.paused",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 422, resp.Code)
	models.DeleteConstraint("spec.paused", gkv)
}

///////////////////////
// Enum array values //
///////////////////////

func TestValidateEndpoint_WithEnumAndIntegerArrayValue_ShouldWork(t *testing.T) {
	gkv := models.GroupKindVersion{
		Group:   "",
		Kind:    "Pod",
		Version: "v1",
	}
	models.DeleteAll()
	enum := []string{"1", "2", "5", "4", "3"}
	var constraint = models.Constraint{
		Enum:             enum,
		Path:             "spec.securityContext.supplementalGroups",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("./resources/validIntegerArray.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 200, resp.Code)
	models.DeleteConstraint("spec.securityContext.supplementalGroups", gkv)

	// Given
	enum = []string{"1", "2", "5", "3"}
	constraint = models.Constraint{
		Enum:             enum,
		Path:             "spec.securityContext.supplementalGroups",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 422, resp.Code)
	models.DeleteConstraint("spec.securityContext.supplementalGroups", gkv)
}

func TestValidateEndpoint_WithEnumAndStringArrayValue_ShouldWork(t *testing.T) {
	gkv := models.GroupKindVersion{
		Group:   "authorization.k8s.io",
		Kind:    "SelfSubjectAccessReview",
		Version: "v1",
	}
	models.DeleteAll()
	enum := []string{"beeyond", "resource", "names", "test", "example"}
	var constraint = models.Constraint{
		Enum:             enum,
		Path:             "status.resourceRules.resourceNames",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp := httptest.NewRecorder()
	c, _ := ioutil.ReadFile("./resources/validStringArray.yaml")
	req, _ := http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 200, resp.Code)
	models.DeleteConstraint("status.resourceRules.resourceNames", gkv)

	// Given
	enum = []string{"beeyond", "names", "test", "example"}
	constraint = models.Constraint{
		Enum:             enum,
		Path:             "status.resourceRules.resourceNames",
		GroupKindVersion: gkv,
	}

	_ = models.SaveConstraint(constraint)

	// When
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/validate", strings.NewReader(string(c)))
	Router.ServeHTTP(resp, req)

	// Then
	assert.Equal(t, 422, resp.Code)
	models.DeleteConstraint("status.resourceRules.resourceNames", gkv)
}

/////////////////////////
// Regex single values //
/////////////////////////

func TestValidateEndpoint_WithRegexAndIntegerValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndStringValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

// TODO: probleme mit casten -> auslassen
func TestValidateEndpoint_WithRegexAndBooleanValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

////////////////////////
// Regex array values //
////////////////////////

func TestValidateEndpoint_WithRegexAndIntegerArrayValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndStringArrayValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithRegexAndBooleanArrayValue_ShouldWork(t *testing.T) {
	// Given
	// When
	// Then
}

func TestValidateEndpoint_WithEmptyYaml_ShouldReturnError(t *testing.T) {
	// Given
	// When
	// Then
}
