package auth

import (
	"anylbapi/internal/database"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AnyLBClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func MakeJWT(user database.User, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AnyLBClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "AnyLB",
			Subject:   user.Username,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		},
		Username: user.Username,
	})

	return token.SignedString([]byte(tokenSecret))
}

func ValidateJWT(tokenString, tokenSecret string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AnyLBClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("incorrect encoding algorithm")
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}

	if claims, ok := token.Claims.(*AnyLBClaims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", fmt.Errorf("invalid token")
}
