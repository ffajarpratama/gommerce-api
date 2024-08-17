package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID string
	Role   string
	jwt.RegisteredClaims
}

const (
	ACCESS_TTL = time.Duration(1 * 24 * time.Hour)
)

func GenerateToken(claims *CustomClaims, secret string) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESS_TTL)),
		Issuer:    "github.com/ffajarpratama",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(secret))
}
