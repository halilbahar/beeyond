package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"kubernetes-validation-beeyond/conf"
	_ "kubernetes-validation-beeyond/docs"
	"kubernetes-validation-beeyond/middleware"
	"net/http"
)

// Creates an Engine with all endpoint, their paths and the used middleware
// Returns: *gin.Engine: an Engine with all defined Endpoints and the used middleware

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	api := router.Group("/api")
	{
		api.GET("/swagger-ui", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "swagger/index.html")
		})
		// validate
		api.POST("/validate", getValidationResult)
		api.Use(middleware.PathSegments())
		api.Use(middleware.ProvideSchema())
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

// Initialises the Router and runs it
func Init() {
	router := GetRouter()
	_ = router.Run(conf.Configuration.Server.HttpPort)
}
