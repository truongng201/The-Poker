package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaimsAccessTokenSub struct {
	Email    string `json:"email"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

// GenerateJWT generates a JWT
func GenerateJWT(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

// ParseJWT parses a JWT
func ParseJWT(tokenString, secret string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, jwt.ErrTokenExpired
		}
		return nil, fmt.Errorf("error parsing JWT: %v", err)
	}

	return token.Claims, nil
}

// GetJWTClaims gets the JWT claims
func GetJWTClaims(tokenString, secret string) (jwt.MapClaims, error) {
	claims, err := ParseJWT(tokenString, secret)

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, jwt.ErrTokenExpired
		}
		return nil, fmt.Errorf("error getting JWT claims: %v", err)
	}

	return claims.(jwt.MapClaims), nil
}
