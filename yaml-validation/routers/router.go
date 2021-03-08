package routers

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"yaml-validation/middleware"
	"yaml-validation/pkg/setting"

	"github.com/gin-gonic/gin"
	_ "yaml-validation/docs"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	api := router.Group("/api")
	{
		// validate
		api.POST("/validate", getValidationResult)
		api.Use(middleware.PathSegments())
		url := ginSwagger.URL("http://localhost:8180/api/swagger/doc.json") // The url pointing to API definition
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

		// constraints
		constraints := api.Group("/constraints")
		{
			constraints.GET("", listRootConstraints)
			constraints.GET("/*path", getConstraintsByPath)

			pathValid := middleware.PathValid()
			constraints.POST("/*path", pathValid, createConstraintByPath)
			constraints.DELETE("/*path", pathValid, deleteConstraintByPath)
			constraints.PATCH("/*path", pathValid, toggleDisableConstraintByPath)
		}
	}

	return router
}

func Init() {
	router := GetRouter()
	_ = router.Run(setting.ServerSetting.HttpPort)
}
