package routers

import (
	"net/http"
	"yaml-validation/models"

	"github.com/gin-gonic/gin"
)

func createConstraint(c *gin.Context) {
	var constraint models.Constraint
	segments := c.GetStringSlice("pathSegments")

	if err := c.ShouldBindJSON(&constraint); err != nil {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if !isValid(&constraint) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var groupKindVersion models.GroupKindVersion
	groupKindVersion, constraint.Path = models.GetGroupKindVersionAndPathFromSegments(segments)

	var lastSegment *string
	if len(segments) != 1 {
		lastSegment = &segments[len(segments)-1]
		segments = segments[0 : len(segments)-1]
	}

	currentSchema, err := models.GetSchemaBySegments(segments)
	// Check if schema was not found or the property was not found. Use the last segment for checking a property
	if err != nil || currentSchema.Properties[*lastSegment] == nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if currentSchema.GroupKindVersion == nil {
		constraint.GroupKindVersion = append(constraint.GroupKindVersion, groupKindVersion)
	} else {
		for _, aGroupKindVersion := range currentSchema.GroupKindVersion {
			constraint.GroupKindVersion = append(constraint.GroupKindVersion, aGroupKindVersion.ToLower())
		}
	}

	models.DeleteConstraint(constraint.Path, &constraint.GroupKindVersion[0])
	if err := models.SaveConstraint(constraint); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
}

func listRootConstraints(c *gin.Context) {
	collection, err := models.GetSchemaCollection()
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var kubernetesRootDefinitions []*models.Schema
	for _, definition := range collection.Schemas {
		groupKindVersions := definition.GroupKindVersion
		if len(groupKindVersions) > 0 && groupKindVersions[0].Kind != "" {
			definition.Constraint = models.GetConstraint("", &groupKindVersions[0])
			kubernetesRootDefinitions = append(kubernetesRootDefinitions, definition)
		}
	}

	c.JSON(http.StatusOK, kubernetesRootDefinitions)
}

func getConstraintsByPath(c *gin.Context) {
	segments := c.GetStringSlice("pathSegments")
	schema, err := models.GetSchemaBySegments(segments)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, schema)
}

func isValid(c *models.Constraint) bool {
	if c.Enum == nil && c.Min == nil && c.Max == nil && c.Regex == "" {
		return false
	}

	return isEnumValid(c) || isMinMaxValid(c) || isRegexValid(c)
}

func isEnumValid(c *models.Constraint) bool {
	return c.Enum != nil && c.Min == nil && c.Max == nil && c.Regex == ""
}

func isMinMaxValid(c *models.Constraint) bool {
	return c.Min != nil && c.Max != nil && c.Enum == nil && c.Regex == ""
}
func isRegexValid(c *models.Constraint) bool {
	return c.Regex != "" && c.Enum == nil && c.Min == nil && c.Max == nil
}

func getAll(c *gin.Context) {
	constr := models.GetConstraints()
	c.JSON(http.StatusOK, constr)
}
