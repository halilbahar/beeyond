package routers

import (
	"net/http"
	"strings"
	"yaml-validation/models"

	"github.com/gin-gonic/gin"
)

func listRootConstraints(c *gin.Context) {
	collection, err := models.GetSchemaCollection()
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var kubernetesRootDefinitions []*models.Schema
	for _, schema := range collection.Schemas {
		groupKindVersions := schema.GroupKindVersion
		if len(groupKindVersions) > 0 {
			schema.Constraint = models.GetConstraint("", groupKindVersions[0])
			kubernetesRootDefinitions = append(kubernetesRootDefinitions, schema)
		}

		for _, property := range schema.Properties {
			var referencePath string
			if property.Reference != "" {
				referencePath = property.Reference
			} else if property.Type == "array" {
				referencePath = property.Items.Reference
			}

			if referencePath != "" {
				split := strings.Split(referencePath, "/")
				definitionName := split[len(split)-1]

				if collection.Schemas[definitionName].Type == "object" && collection.Schemas[definitionName].Properties != nil {
					property.IsKubernetesObject = true
				}
			}
		}
	}

	c.JSON(http.StatusOK, kubernetesRootDefinitions)
}

func getConstraintsByPath(c *gin.Context) {
	segments := c.GetStringSlice("pathSegments")
	schema, err := models.GetSchemaBySegments(segments)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, schema)
}

func createConstraintByPath(c *gin.Context) {
	var constraint models.Constraint
	if err := c.ShouldBindJSON(&constraint); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	segments := c.GetStringSlice("pathSegments")
	var lastSegment string
	if len(segments) != 1 {
		lastSegment = segments[len(segments)-1]
		segments = segments[0 : len(segments)-1]
	}

	schema, _ := models.GetSchemaBySegments(segments)

	if schema.Properties[lastSegment] != nil && !schema.Properties[lastSegment].IsKubernetesObject && !constraint.IsValid(schema.Properties[lastSegment].Type) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	groupKindVersionInterface, _ := c.Get("groupKindVersion")
	constraint.Path = c.GetString("propertyPath")

	// check if constraint on apiVersion or kind
	if strings.HasSuffix(constraint.Path, "apiVersion") || strings.HasSuffix(constraint.Path, "kind") {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	constraint.GroupKindVersion = groupKindVersionInterface.(models.GroupKindVersion)

	models.DeleteConstraint(constraint.Path, constraint.GroupKindVersion)
	if err := models.SaveConstraint(constraint); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
}

func deleteConstraintByPath(c *gin.Context) {
	groupKindVersion, _ := c.Get("groupKindVersion")
	propertyPath := c.GetString("propertyPath")

	models.DeleteConstraint(propertyPath, groupKindVersion.(models.GroupKindVersion))
	c.Writer.WriteHeader(http.StatusNoContent)
}

func toggleDisableConstraintByPath(c *gin.Context) {
	groupKindVersionInterface, _ := c.Get("groupKindVersion")
	propertyPath := c.GetString("propertyPath")
	groupKindVersion := groupKindVersionInterface.(models.GroupKindVersion)

	constraint := models.GetConstraint(propertyPath, groupKindVersion)
	// If the no constraint exits and the user wants to disable the path, create a new constraint
	if constraint == nil {
		constraint = &models.Constraint{
			Path:             propertyPath,
			Disabled:         false,
			GroupKindVersion: groupKindVersion,
		}
	}

	constraint.Disabled = !constraint.Disabled
	models.DeleteConstraint(propertyPath, groupKindVersion)
	if models.SaveConstraint(*constraint) != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
