package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	var kubernetesRootDefinitions []models.Schema
	for _, definition := range collection.Schema {
		groupKindVersions := definition.GroupKindVersion
		if len(groupKindVersions) > 0 && groupKindVersions[0].Kind != "" {
			kubernetesRootDefinitions = append(kubernetesRootDefinitions, definition)
		}
	}

	// TODO: add constraint from the database to the definition if one is present

	c.JSON(http.StatusOK, kubernetesRootDefinitions)
}
