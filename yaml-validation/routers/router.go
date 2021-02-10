package routers

import (
	"yaml-validation/middleware"
	"yaml-validation/pkg/setting"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
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

	return router
}

func Init() {
	router := GetRouter()
	_ = router.Run(setting.ServerSetting.HttpPort)
}
