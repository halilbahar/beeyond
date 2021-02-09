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

	if !isValid(&constraint) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var gkv models.GroupKindVersion
	gkv, constraint.Path = getGroupKindVersionAndPathFromPath(segments)

	collection, err := models.GetSchemaCollection()
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var currentSchema *models.Schema

	for i, segment := range segments[0 : len(segments)-1] {
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

			// If the specified path of the user does not exist, break
			if referencePath == "" {
				break
			}

			// We want the last part of the reference
			// Example: #/definitions/io.k8s.api.apps.v1.DeploymentSpec
			split := strings.Split(referencePath, "/")
			definitionName := split[len(split)-1]

			currentSchema = collection.Schemas[definitionName]
		}
	}

	if currentSchema.GroupKindVersion == nil {
		constraint.GroupKindVersion = append(constraint.GroupKindVersion, gkv)
	} else {
		for _, gkvL := range currentSchema.GroupKindVersion {
			gkvLower := models.GkvToLower(&gkvL)

			constraint.GroupKindVersion = append(constraint.GroupKindVersion, *gkvLower)
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
		if constraintPath == "" {
			prop.Constraint = models.GetConstraint(k, &gkv)
		} else {
			prop.Constraint = models.GetConstraint(constraintPath+"."+k, &gkv)
		}
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
