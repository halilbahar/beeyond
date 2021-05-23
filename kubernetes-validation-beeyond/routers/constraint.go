package routers

import (
	"kubernetes-validation-beeyond/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Adds all root elements with their constraints
// to the body of the current http request
// Parameter: c *gin.Context
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

// Adds the schema according to the current path
// with its constraints to the body of the current constraint
// Parameter: c (*gin.Context): Contains the path to the schema
// Possible status codes:
// 		- 404, if the path was not valid (doesn't exist)
// 		- 200, if schema was found
func getConstraintsByPath(c *gin.Context) {
	segments := c.GetStringSlice("pathSegments")
	schema, err := models.GetSchemaBySegments(segments)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, schema)
}

// Saves the given constraint, only if the given
// exists and is valid for constraints.
// Parameter: c (*gin.Context): Contains the parameters
// that were sent (path and constraint)
// Possible status Codes:
// 		- 201, if constraint was saved
// 		- 400, invalid constraint / path
// 		- 500, if constraint / path valid, but saving in database didn't work
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

// Deletes the constraint with the given path
// Parameter: c (*gin.Context): contains the path of
// the constraint we want to delete
// Possible status codes:
// 		- 400, if path was not found
// 		- 204, if the deletion was successful
func deleteConstraintByPath(c *gin.Context) {
	groupKindVersion, _ := c.Get("groupKindVersion")
	propertyPath := c.GetString("propertyPath")

	if models.DeleteConstraint(propertyPath, groupKindVersion.(models.GroupKindVersion)).DeletedCount == 0 {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

// Toggles the "disabled" from the constraint with the given path
// If the given constraint does not exist, it will be created
// Parameter: c (*gin.Context): Contains the path to the constraint
// where we want to toggle the "disable" field
// Possible status codes:
// 		- 200, if toggle worked
// 		- 400, if field is required
// 		- 500, if problems with the database-connection occur
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
