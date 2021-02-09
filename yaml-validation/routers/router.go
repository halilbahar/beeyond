package routers

import (
	"yaml-validation/middleware"
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
		constraints := api.Group("/constraints")
		{
			constraints.Use(middleware.KubernetesPath())
			constraints.GET("", listRootConstraints)
			constraints.POST("/*path", createConstraint)
			constraints.GET("/*path", getConstraintsByPath)

		}
		api.GET("/constraintsall/", getAll)
	}

	_ = router.Run(setting.ServerSetting.HttpPort)
}
