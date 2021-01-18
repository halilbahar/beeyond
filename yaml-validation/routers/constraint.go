package routers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addConstraint(c *gin.Context) {
	var constraint models.Constraint
	if err := c.ShouldBindJSON(&constraint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if models.SaveConstraintToDb(constraint) != nil {
		c.Writer.WriteHeader(400)
		return
	}

	c.Writer.WriteHeader(201)
}

func listConstraints(c *gin.Context) {
	c.JSON(200, models.GetConstraints())
}
