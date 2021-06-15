package middleware
//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/coreos/go-oidc/v3/oidc"
//	"golang.org/x/oauth2"
//	"net/http"
//)
//
//func Oidc(ctx *gin.Context) gin.HandlerFunc {
//
//	provider, err := oidc.NewProvider(ctx, "localhost:8280")
//	if err != nil {
//		// handle error
//	}
//
//	// Configure an OpenID Connect aware OAuth2 client.
//	oauth2Config := oauth2.Config{
//		ClientID:     clientID,
//		ClientSecret: clientSecret,
//		RedirectURL:  redirectURL,
//
//		// Discovery returns the OAuth2 endpoints.
//		Endpoint: provider.Endpoint(),
//
//		// "openid" is a required scope for OpenID Connect flows.
//		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
//	}
//	var verifier = provider.Verifier(&oidc.Config{ClientID: clientID})
//
//	return func (w http.ResponseWriter, r *http.Request) {
//		// Verify state and errors.
//
//		oauth2Token, err := oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
//		if err != nil {
//			// handle error
//		}
//
//		// Extract the ID Token from OAuth2 token.
//		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
//		if !ok {
//			// handle missing token
//		}
//
//		// Parse and verify ID Token payload.
//		idToken, err := verifier.Verify(ctx, rawIDToken)
//		if err != nil {
//			// handle error
//		}
//
//		// Extract custom claims
//		var claims struct {
//			Email    string `json:"email"`
//			Verified bool   `json:"email_verified"`
//		}
//		if err := idToken.Claims(&claims); err != nil {
//			// handle error
//		}
//	}
//}
