package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"yaml-validation/models"
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
	fmt.Print("hello")
	request, _ := http.NewRequest("POST", "/api/constraints/RuntimeClassList-node.k8s.io-v1beta1/metadata/continue", bytes.NewBuffer(b))
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
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
	request, _ = http.NewRequest("PATCH", "/api/constraints/Deployment-apps-v1/spec/template/metadata/clusterName", bytes.NewBuffer(b))
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
	Router.ServeHTTP(responseRecorder, request)

	// Then
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestConstraintCreation_EmptyBody_Fail(t *testing.T) {
	// Given When
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/constraints/Deployment-apps-v1/spec/replicas", nil)
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
