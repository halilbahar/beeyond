package routers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/validate", models.GetValidationResult)
	r.POST("/api/constraints", models.AddConstraint)
	return r
}