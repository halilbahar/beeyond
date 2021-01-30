package routers

import (
	"yaml-validation/pkg/setting"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	api := router.Group("/api")
	{
		// validate
		api.POST("/validate", getValidationResult)

		// constraints
		api.GET("/constraints", listRootConstraints)
		api.POST("/constraints", createConstraint)
		api.GET("/constraints/*path", getConstraintsByPath)
	}

	_ = router.Run(setting.ServerSetting.HttpPort)
}
