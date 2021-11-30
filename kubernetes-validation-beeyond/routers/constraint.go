package routers

import (
	"kubernetes-validation-beeyond/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary Find root constraints
// @Description Finds all root schemes and their constraints
// @Tags Constraint
// @Success 200 {string} string	"ok"
// @Failure 500 {string} string "internal server error"
// @Router /api/constraints/ [get]
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

		delete(schema.Properties, "apiVersion")
		delete(schema.Properties, "kind")

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

// @Summary Find constraints by path
// @Description Finds the schema and its constraints according to the given path
// @Tags Constraint
// @Param  "path"     path    string     true        "path"
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "bad request"
// @Router /api/constraints/{path} [get]
func getConstraintsByPath(c *gin.Context) {
	segments := c.GetStringSlice("pathSegments")
	schema, err := models.GetSchemaBySegments(segments)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, schema)
}

// @Summary Creates a new constraint
// @Description creates a new constraint and adds it to the database. If the constraint already exists it gets replaced.
// @Tags Constraint
// @Accept  json
// @Param  "path"     path    string     true        "path"
// @Success 201 {string} string	"created"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /api/constraints/{path} [post]
func createConstraintByPath(c *gin.Context) {
	var constraint models.Constraint
	if err := c.ShouldBindJSON(&constraint); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	lastSegment := c.GetString("lastPropertyName")
	schemaInterface, _ := c.Get("schema")
	schema := schemaInterface.(*models.Schema)

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

// @Summary Delete constraint
// @Description Deletes the constraint with the given path
// @Tags Constraint
// @Param  "path"     path    string     true        "path"
// @Success 204 {string} string	"no content"
// @Failure 400 {string} string "bad request"
// @Router /api/constraints/{path} [delete]
func deleteConstraintByPath(c *gin.Context) {
	groupKindVersion, _ := c.Get("groupKindVersion")
	propertyPath := c.GetString("propertyPath")

	if models.DeleteConstraint(propertyPath, groupKindVersion.(models.GroupKindVersion)).DeletedCount == 0 {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

// @Summary Toggle disabled on constraint
// @Description Toggles the "disabled" from the constraint with the given path. If the given constraint does not exist, it will be created
// @Tags Constraint
// @Param  "path"     path    string     true        "path"
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /api/constraints/{path} [patch]
func toggleDisableConstraintByPath(c *gin.Context) {
	groupKindVersionInterface, _ := c.Get("groupKindVersion")
	propertyPath := c.GetString("propertyPath")
	groupKindVersion := groupKindVersionInterface.(models.GroupKindVersion)

	constraint := models.GetConstraint(propertyPath, groupKindVersion)
	// If no constraint exits and the user wants to disable the path, create a new constraint
	if constraint == nil {
		constraint = &models.Constraint{
			Path:             propertyPath,
			Disabled:         false,
			GroupKindVersion: groupKindVersion,
		}
	}

	lastSegment := c.GetString("lastPropertyName")
	schemaInterface, _ := c.Get("schema")
	schema := schemaInterface.(*models.Schema)

	for _, req := range schema.Required {
		if req == lastSegment {
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	constraint.Disabled = !constraint.Disabled
	models.DeleteConstraint(propertyPath, groupKindVersion)
	if models.SaveConstraint(*constraint) != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	schema.Constraint = constraint
	c.Writer.WriteHeader(http.StatusOK)
}
