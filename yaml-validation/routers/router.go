package routers

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	api := router.Group("/api")
	{
		// validate
		api.POST("/validate", getValidationResult)

		// constraints
		api.GET("/constraints", listConstraints)
		api.POST("/constraints", createConstraint)
	}

	_ = router.Run(":8180")
}
