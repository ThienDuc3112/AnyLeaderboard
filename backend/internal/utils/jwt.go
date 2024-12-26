package utils

import (
	"anylbapi/internal/database"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func MakeAccessTokenJWT(user database.User, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessTokenClaims{
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
	TokenID         int32 `json:"asdf"`
	RotationCounter int32 `json:"poiu"`
}

func MakeRefreshTokenJWT(refreshToken database.RefreshToken, tokenSecret string, expires_at time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "AnyLB",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expires_at),
		},
		TokenID:         refreshToken.ID,
		RotationCounter: refreshToken.RotationCounter,
	})

	return token.SignedString([]byte(tokenSecret))
}

func ValidateAccessToken(tokenString string, tokenSecret string) (*AccessTokenClaims, error) {
	// Parse the token with a key function
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func ValidateRefreshToken(tokenString string, tokenSecret string) (*RefreshTokenClaims, error) {
	// Parse the token with a key function
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	if claims, ok := token.Claims.(*RefreshTokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
