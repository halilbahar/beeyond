package routers

import (
	"net/http"
	"strings"
	"yaml-validation/models"

	"github.com/gin-gonic/gin"
)

func createConstraintByPath(c *gin.Context) {
	var constraint models.Constraint
	segments := c.GetStringSlice("pathSegments")

	if err := c.ShouldBindJSON(&constraint); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the constraint is valid (enum or min-max or regex based on type)
	if !constraint.IsValid() {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if the path exists for kubernetes
	if !models.IsValidConstraintPath(segments) {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	var groupKindVersion models.GroupKindVersion
	groupKindVersion, constraint.Path = models.GetGroupKindVersionAndPathFromSegments(segments)
	constraint.GroupKindVersion = append(constraint.GroupKindVersion, groupKindVersion)

	models.DeleteConstraint(constraint.Path, constraint.GroupKindVersion[0])
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
			definition.Constraint = models.GetConstraint("", groupKindVersions[0])
			kubernetesRootDefinitions = append(kubernetesRootDefinitions, definition)
		}
		for _, property := range definition.Properties {
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

func toggleDisableConstraintByPath(c *gin.Context) {
	segments := c.GetStringSlice("pathSegments")

	// check if the path exists for kubernetes
	if !models.IsValidConstraintPath(segments) {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	// Get the group kind and version and fetch the constraint from the database with that information
	groupKindVersion, path := models.GetGroupKindVersionAndPathFromSegments(segments)
	constraint := models.GetConstraint(path, groupKindVersion)
	// If the no constraint exits and the user wants to disable the path, create a new constraint
	if constraint == nil {
		constraint = &models.Constraint{
			Path:             path,
			Disabled:         false,
			GroupKindVersion: []models.GroupKindVersion{groupKindVersion},
		}
	}

	constraint.Disabled = !constraint.Disabled
	models.DeleteConstraint(path, groupKindVersion)
	if models.SaveConstraint(*constraint) != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
