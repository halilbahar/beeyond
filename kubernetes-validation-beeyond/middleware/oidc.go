package middleware

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

func Oidc(r *http.Request) gin.HandlerFunc {
	return func (ctx *gin.Context) {
		provider, err := oidc.NewProvider(ctx, "https://localhost:8280/auth/realms/school")
		if err != nil {
			// handle error
		}

		// Configure an OpenID Connect aware OAuth2 client.
		oauth2Config := oauth2.Config{
			ClientID: "beeyond-spa",

			// Discovery returns the OAuth2 endpoints.
			Endpoint: provider.Endpoint(),

			// "openid" is a required scope for OpenID Connect flows.
			Scopes: []string{oidc.ScopeOpenID, "offline_access"},
		}
		var verifier = provider.Verifier(&oidc.Config{ClientID: "beeyond-spa"})
		// Verify state and errors.

		oauth2Token, err := oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			// handle error
		}

		// Extract the ID Token from OAuth2 token.
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			// handle missing token
		}

		// Parse and verify ID Token payload.
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			// handle error
		}

		// Extract custom claims
		var claims struct {
			Email    string `json:"email"`
			Verified bool   `json:"email_verified"`
		}
		if err := idToken.Claims(&claims); err != nil {
			// handle error
		}
	}
}
