package middleware

import (
	"github.com/gin-gonic/gin"
	"yaml-validation/models"
)

func ProvideSchema() gin.HandlerFunc {
	return func(c *gin.Context) {
		segments := c.GetStringSlice("pathSegments")
		var lastSegment string
		if len(segments) != 1 {
			lastSegment = segments[len(segments)-1]
			segments = segments[0 : len(segments)-1]
		}

		schema, _ := models.GetSchemaBySegments(segments)

		c.Set("schema", schema)
		c.Set("lastPropertyName", lastSegment)
		c.Next()
	}
}
