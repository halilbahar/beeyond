package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Oidc() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		authHeader := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			ctx.Writer.WriteHeader(http.StatusUnauthorized)
			ctx.Writer.Write([]byte("Malformed Token"))
		} else {

			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(SECRETKEY), nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx.Set("props", claims)
				ctx.Next()
			} else {
				fmt.Println(err)
				ctx.Writer.WriteHeader(http.StatusUnauthorized)
				ctx.Writer.Write([]byte("Unauthorized"))
			}
		}
	}
}
