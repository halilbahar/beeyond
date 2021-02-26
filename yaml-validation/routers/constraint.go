package routers

import (
	"net/http"
	"strings"
	"yaml-validation/models"

	"github.com/gin-gonic/gin"
)

//
// @Summary Creates a new constraint
// @Description creates a new constraint and adds it to the database. If the constraint already exists it gets replaced.
// @Accept  json
// @Param  "path"     path    string     true        "path"
// @Success 201 {string} string	"created"
// @Failure 400 {string} string "bad request"
// @Router /api/constraints/{path} [post]
func createConstraint(c *gin.Context) {
	var constraint models.Constraint
	segments := c.GetStringSlice("pathSegments")

	if err := c.ShouldBindJSON(&constraint); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if !constraint.IsValid() {
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
		for _, property := range definition.Properties {
			var referencePath string
			if property.Reference != "" {
				referencePath = property.Reference
			} else if property.Type == "array"{
				referencePath = property.Items.Reference
			}

			if referencePath != "" {
				split := strings.Split(referencePath, "/")
				definitionName := split[len(split)-1]

				if collection.Schemas[definitionName].Type == "object" && collection.Schemas[definitionName].Properties != nil{
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
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, schema)
}

func getAll(c *gin.Context) {
	constr := models.GetConstraints()
	c.JSON(http.StatusOK, constr)
}

func disableConstraint(c *gin.Context) {
	var constraint *models.Constraint

	gkv, path := models.GetGroupKindVersionAndPathFromSegments(c.GetStringSlice("pathSegments"))

	if constraint = models.GetConstraint(path, &gkv); constraint == nil{
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	constraint.Disabled = true
	models.DeleteConstraint(constraint.Path, &constraint.GroupKindVersion[0])
	models.SaveConstraint(*constraint)

	c.Writer.WriteHeader(http.StatusOK)
}