package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func KubernetesPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Param("path")
		// Simplify path so it is easier to split and find the object
		// /deployment-apps-v1/metadata/ -> deployment-apps-v1/metadata
		trimmedPath := strings.Trim(path, "/")
		segments := strings.Split(trimmedPath, "/")

		c.Set("pathSegments", segments)
		c.Next()
	}
}
