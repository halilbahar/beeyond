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

func listConstraints(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetConstraints())
}
