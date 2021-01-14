package models

import (
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddConstraint(c *gin.Context) {
	var constraint  services.Constraint
	if err := c.ShouldBindJSON(&constraint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if services.SaveConstraint(constraint) != nil {
		c.Writer.WriteHeader(400)
		return
	}

	c.Writer.WriteHeader(201)
}

