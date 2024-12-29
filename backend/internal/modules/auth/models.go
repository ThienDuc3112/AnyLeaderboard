package auth

import (
	"anylbapi/internal/database"
	"errors"
)

// ============ Constants ============
const (
	cookieKeyRefreshToken = "refresh_token"
)

// ============ Request body type ============
type loginReqBody struct {
	Username string `json:"username" validate:"required,isSafeName"`
	Password string `json:"password" validate:"required"`
}

type signUpReqBody struct {
	Username    string `json:"username" validate:"required,min=3,max=64,isUsername"`
	DisplayName string `json:"displayName" validate:"required,min=3,max=64,isSafeName"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,max=64"`
}

// ============ Service param and return types ============
type loginParam struct {
	loginReqBody
	DeviceInfo string
	IpAddress  string
}

type refreshParam struct {
	RefreshToken string
	DeviceInfo   string
	IpAddress    string
}

type tokensReturn struct {
	accessToken     string
	refreshToken    string
	refreshTokenRaw database.RefreshToken
}

type signUpParam struct {
	signUpReqBody
}

// ============ Service returned errors ============
var (
	errNoUser                  = errors.New("user don't exist")
	errIncorrectPassword       = errors.New("incorrect password")
	errUsernameTaken           = errors.New("username is taken")
	errEmailUsed               = errors.New("email is already used")
	errNoTokenExist            = errors.New("refresh token don't exist")
	errMismatchRotationCounter = errors.New("refresh token rotationCounter don't match")
	errMismatchUpdatedRC       = errors.New("refresh token rotationCounter don't match after update, potential race condition")
	errTokenRevoked            = errors.New("refresh token revoked")
	errInvalidToken            = errors.New("invalid token")
)
