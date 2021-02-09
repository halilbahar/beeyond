package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"yaml-validation/models"
)

func createConstraint(c *gin.Context) {
	var constraint models.Constraint

	segments := c.GetStringSlice("pathSegments")

	if err := c.ShouldBindJSON(&constraint); err != nil {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if isInValid(&constraint) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var gkv models.GroupKindVersion
	gkv, constraint.Path = getGroupKindVersionAndPathFromPath(segments)

	constraint.GroupKindVersion = append(constraint.GroupKindVersion, gkv)

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
			kubernetesRootDefinitions = append(kubernetesRootDefinitions, definition)

		}
	}

	// TODO: add constraint from the database to the definition if one is present

	c.JSON(http.StatusOK, kubernetesRootDefinitions)
}

func getConstraintsByPath(c *gin.Context) {
	segments := c.GetStringSlice("pathSegments")

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

			var referencePath string
			if property.Reference != "" {
				referencePath = property.Reference
			} else if property.Items != nil {
				referencePath = property.Items.Reference
			}

			// If the specified path of the user does not exist, return
			// This means the user requested something other than object (string, int, ...)
			if referencePath == "" {
				c.Writer.WriteHeader(http.StatusBadRequest)
				return
			}

			// We want the last part of the reference
			// Example: #/definitions/io.k8s.api.apps.v1.DeploymentSpec
			split := strings.Split(referencePath, "/")
			definitionName := split[len(split)-1]

			currentSchema = collection.Schemas[definitionName]
		}
	}

	gkv, constraintPath := getGroupKindVersionAndPathFromPath(segments)

	for k, prop := range currentSchema.Properties {
		prop.Constraint = models.GetConstraint(constraintPath+"."+k, &gkv)
	}

	c.JSON(http.StatusOK, currentSchema)
}

func getGroupKindVersionAndPathFromPath(segments []string) (models.GroupKindVersion, string) {
	var gkv models.GroupKindVersion
	parts := strings.Split(segments[0], "-")
	gkv.Kind = parts[0]
	if len(parts) == 3 {
		gkv.Group = parts[1]
		gkv.Version = parts[2]
	} else {
		gkv.Version = parts[1]
	}

	constraintPath := strings.Join(segments[1:], ".")
	return gkv, constraintPath
}

func isInValid(c *models.Constraint) bool {
	return c.Enum != nil && (c.Min != nil || c.Max != nil || c.Regex != "") ||
		(c.Min != nil || c.Max != nil) && (c.Enum != nil || c.Regex != "") ||
		c.Regex != "" && (c.Min != nil || c.Max != nil || c.Enum != nil)

}

func getAll(c *gin.Context) {
	constr := models.GetConstraints()
	c.JSON(http.StatusOK, constr)
}
