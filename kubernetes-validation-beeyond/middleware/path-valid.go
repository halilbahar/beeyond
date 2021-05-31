package middleware

import (
	"github.com/gin-gonic/gin"
	"kubernetes-validation-beeyond/models"
	"net/http"
)

// Middleware which checks weather the entered path is valid
// and sets the field groupKindVersion and propertyPath if the path
// is valid
// if the path is not valid the method will return 404 NotFound
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
