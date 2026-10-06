package auth

import (
	"time"

	"github.com/Chefmine8/OpenCroupier/internal"
	"github.com/Chefmine8/OpenCroupier/internal/env"
	"github.com/golang-jwt/jwt/v5"
)

func newClaims(uid string, userName string) *internal.Claims {
	expirationTime := time.Now().Add(6 * time.Hour)
	claims := &internal.Claims{
		UserUID:  uid,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return claims
}

func NewJWT(uid string, userName string) (string, error) {
	claims := newClaims(uid, userName)

	JwtKey := env.GetEnvVariable("JWTKEY", "ERROR")
	if JwtKey == "ERROR" || JwtKey == "" {
		panic("JWTKEY ERROR")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JwtKey))
}
