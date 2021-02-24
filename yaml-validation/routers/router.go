package routers

import (
	"yaml-validation/middleware"
	"yaml-validation/pkg/setting"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	api := router.Group("/api")
	{
		// validate
		api.POST("/validate", getValidationResult)
		api.Use(middleware.KubernetesPath())

		// constraints
		constraints := api.Group("/constraints")
		{
			constraints.GET("", listRootConstraints)
			constraints.GET("/*path", getConstraintsByPath)
			constraints.POST("/*path", createConstraint)
		}
		api.POST("/disable-constraint/*path", disableConstraint)
		api.GET("/constraintsall/", getAll)
	}

	return router
}

func Init() {
	router := GetRouter()
	_ = router.Run(setting.ServerSetting.HttpPort)
}
