package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yaml-validation/models"
)

func PathValid() gin.HandlerFunc {
	return func(c *gin.Context) {
		segments := c.GetStringSlice("pathSegments")
		if !models.IsValidConstraintPath(segments) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		groupKindVersion, propertyPath := models.GetGroupKindVersionAndPathFromSegments(segments)
		c.Set("groupKindVersion", groupKindVersion)
		c.Set("propertyPath", propertyPath)
		c.Next()
	}
}
