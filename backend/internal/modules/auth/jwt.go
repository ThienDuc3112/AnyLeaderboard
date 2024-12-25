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

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	ID               int32 `json:"asdf"`
	Rotation_counter int32 `json:"poiu"`
}

func MakeRefreshTokenJWT(refreshToken database.RefreshToken, tokenSecret string, expires_at time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "AnyLB",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expires_at),
		},
		ID:               refreshToken.ID,
		Rotation_counter: refreshToken.RotationCounter,
	})

	return token.SignedString([]byte(tokenSecret))
}

func ValidateToken[T jwt.Claims](tokenString string, tokenSecret string) (*T, error) {
	var claims T
	// Parse the token with a key function
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is as expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract and validate the claims
	if claims, ok := token.Claims.(T); ok && token.Valid {
		return &claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
