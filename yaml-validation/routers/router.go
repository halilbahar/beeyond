package routers

import (
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
		api.GET("/constraints", listConstraints)
		api.POST("/constraints", createConstraint)
	}

	return router
}

func Init() {
	router := GetRouter()
	_ = router.Run(setting.ServerSetting.HttpPort)
}
