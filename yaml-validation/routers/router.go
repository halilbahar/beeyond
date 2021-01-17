package routers

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.POST("/api/validate", getValidationResult)
	r.POST("/api/constraints", addConstraint)
	r.GET("api/constraints", listConstraints)
	r.Run(":8180")
}
