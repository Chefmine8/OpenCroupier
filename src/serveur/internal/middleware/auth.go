package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Chefmine8/OpenCroupier/internal"
	"github.com/Chefmine8/OpenCroupier/internal/env"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No AuthHeader", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			http.Error(w, "Invalid AuthHeader", http.StatusUnauthorized)
			return
		}

		claims := &internal.Claims{}
		token, err := jwt.ParseWithClaims(authHeaderParts[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetEnvVariable("JWTKEY", "ERROR")), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			http.Error(w, fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]), http.StatusBadRequest)
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "UID", claims.UserUID)
		ctx = context.WithValue(ctx, "UserName", claims.UserName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
