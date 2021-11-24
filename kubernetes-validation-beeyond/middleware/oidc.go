package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"kubernetes-validation-beeyond/conf"
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
			keytxt := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAg97iKoEpqlBCa8bNSyTWj42JPUFIaORJoEadppCNZWhiy7hE5eUxRX+kvsREGFNeZS8LnQoX1pmXaLq9ZBSt3W8VtlEQRt6z9atHaraaE/zzR8Y0RjD60QcJT9TisAK+Ju/NRIkfUkZ4jPaBXUMfTKmbBqQkef/DVJRHJh0NVg9DZ1P7t01sFp7MxDYW+6m0hmSoHuZER3URFvKpaNShowULBbwme0h48j81t9148ah6hUZgv8uAX3Op3fYxWgWRobMTDLLaKUtZYbmfe/RHWp7u4BR6GzpjeKyFF5ugYhjEfTGmKdDj8cljoTpqRt9MpwS8KazRkjB5bjfQ/zgJFQIDAQAB
-----END PUBLIC KEY-----`
			print(keytxt)
			key := conf.ConvertStringToRSA()
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return key, nil
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
