package routers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

func getValidationResult(c *gin.Context) {
	var json models.Content
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	c.Writer.WriteHeader(200)
}
