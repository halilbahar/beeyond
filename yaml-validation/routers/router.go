package routers

import (
	"../pkg/setting"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.POST("/api/validate", getValidationResult)
	r.POST("/api/constraints", addConstraint)
	r.GET("api/constraints", listConstraints)
	r.Run(setting.ServerSetting.HttpPort)
}
