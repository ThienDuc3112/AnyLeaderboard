package auth

import (
	"anylbapi/internal/models"
	"errors"
)

// ============ Service param and return types ============
type LoginParam struct {
	Username   string
	Password   string
	DeviceInfo string
	IpAddress  string
}

type RefreshParam struct {
	RefreshToken string
	DeviceInfo   string
	IpAddress    string
}

type LoginsReturn struct {
	AccessToken     string
	RefreshToken    string
	RefreshTokenRaw models.RefreshToken
	User            models.UserPreview
}

type SignUpParam struct {
	Username    string
	DisplayName string
	Email       string
	Password    string
}

// ============ Service returned errors ============
var (
	ErrNoUser                  = errors.New("user don't exist")
	ErrIncorrectPassword       = errors.New("incorrect password")
	ErrUsernameTaken           = errors.New("username is taken")
	ErrEmailUsed               = errors.New("email is already used")
	ErrNoTokenExist            = errors.New("refresh token don't exist")
	ErrMismatchRotationCounter = errors.New("refresh token rotationCounter don't match")
	ErrMismatchUpdatedRC       = errors.New("refresh token rotationCounter don't match after update, potential race condition")
	ErrTokenRevoked            = errors.New("refresh token revoked")
	ErrInvalidToken            = errors.New("invalid token")
)
