package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Rbac() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.Keys["props"].(jwt.MapClaims)
		tmp := claims["realm_access"].(map[string]interface{})
		role := tmp["roles"].([]interface{})[0].(string)

		if role != "teacher" {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
