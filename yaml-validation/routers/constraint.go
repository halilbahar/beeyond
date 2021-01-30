package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"yaml-validation/models"
)

func createConstraint(c *gin.Context) {
	var constraint models.Constraint

	if err := c.ShouldBindJSON(&constraint); err != nil {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

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
			kubernetesRootDefinitions = append(kubernetesRootDefinitions, definition)
		}
	}

	// TODO: add constraint from the database to the definition if one is present

	c.JSON(http.StatusOK, kubernetesRootDefinitions)
}

func getConstraintsByPath(c *gin.Context) {
	path := c.Param("path")
	// Simplify path so it is easier to split and find the object
	// /deployment-apps-v1/metadata/ -> deployment-apps-v1/metadata
	trimmedPath := strings.Trim(path, "/")
	segments := strings.Split(trimmedPath, "/")

	collection, err := models.GetSchemaCollection()
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var currentSchema *models.Schema

	for i, segment := range segments {
		// On the first element search for the GroupKindVersion
		if i == 0 {
		schemaLoop:
			for _, schema := range collection.Schemas {

				for _, groupKindVersion := range schema.GroupKindVersion {
					var group string
					if groupKindVersion.Group == "" {
						group = ""
					} else {
						group = "-" + groupKindVersion.Group
					}

					groupKindVersionString := groupKindVersion.Kind + group + "-" + groupKindVersion.Version
					if strings.ToLower(segment) == strings.ToLower(groupKindVersionString) {
						currentSchema = schema
						break schemaLoop
					}
				}
			}

			if currentSchema == nil {
				c.Writer.WriteHeader(http.StatusBadRequest)
				return
			}
		} else {
			property := currentSchema.Properties[segment]

			if property == nil {
				c.Writer.WriteHeader(http.StatusBadRequest)
				return
			}

			// We want the last part of the reference
			// Example: #/definitions/io.k8s.api.apps.v1.DeploymentSpec
			split := strings.Split(property.Reference, "/")
			definitionName := split[len(split)-1]

			schema := collection.Schemas[definitionName]
			if schema == nil {
				c.Writer.WriteHeader(http.StatusBadRequest)
				return
			}

			currentSchema = schema
		}
	}

	// TODO: Add the constraints to currentSchema

	c.JSON(http.StatusOK, currentSchema)
}
