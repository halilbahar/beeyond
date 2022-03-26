package middleware

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"kubernetes-validation-beeyond/conf"
	"net/http"
	"strings"
)

func Oidc() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			keyRSA := FetchKeycloakPublicKey()
			token, err := ParseJwt(jwtToken, keyRSA)
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("props", claims)
				c.Next()
			} else {
				fmt.Println(err)
				c.Writer.WriteHeader(http.StatusUnauthorized)

				c.Writer.Write([]byte("Unauthorized"))
			}
		}
	}
}

func ParseJwt(jwtToken string, keyRSA *rsa.PublicKey) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return keyRSA, nil
	})
	return token, err
}

func FetchKeycloakPublicKey() *rsa.PublicKey {
	var keyRaw map[string]interface{}
	resp, _ := http.Get(conf.Configuration.Authentication.Url + "/auth/realms/" + conf.Configuration.Authentication.Realm)
	json.NewDecoder(resp.Body).Decode(&keyRaw)
	key := "-----BEGIN PUBLIC KEY-----\n" + keyRaw["public_key"].(string) + "\n-----END PUBLIC KEY-----"
	keyRSA := conf.ConvertStringToRSA(key)

	return keyRSA
}
