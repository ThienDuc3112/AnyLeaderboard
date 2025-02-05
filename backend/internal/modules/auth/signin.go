package auth

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/utils"
	"context"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func (s AuthService) Login(context context.Context, param LoginParam) (TokensReturn, error) {
	loginWithEmail := false
	if strings.Contains(param.Username, "@") {
		loginWithEmail = true
	}

	var user database.User
	var err error
	if loginWithEmail {
		user, err = s.repo.GetUserByEmail(context, param.Username)
	} else {
		user, err = s.repo.GetUserByUsername(context, param.Username)
	}

	if err == pgx.ErrNoRows {
		return TokensReturn{}, ErrNoUser
	} else if err != nil {
		return TokensReturn{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))
	if err != nil {
		return TokensReturn{}, ErrIncorrectPassword
	}

	accessToken, err := utils.MakeAccessTokenJWT(user, os.Getenv(constants.EnvKeySecret), constants.AccessTokenDuration)
	if err != nil {
		return TokensReturn{}, err
	}

	refreshTokenParam := database.CreateNewRefreshTokenParams{
		UserID: user.ID,
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(constants.RefreshTokenDuration),
			Valid: true,
		},
		DeviceInfo: param.DeviceInfo,
		IpAddress:  param.IpAddress,
	}

	rToken, err := s.repo.CreateNewRefreshToken(context, refreshTokenParam)
	if err != nil {
		return TokensReturn{}, err
	}

	refreshTokenStr, err := utils.MakeRefreshTokenJWT(rToken, os.Getenv(constants.EnvKeySecret), rToken.ExpiresAt.Time)
	if err != nil {
		return TokensReturn{}, err
	}

	rt := models.RefreshToken{
		ID:              rToken.ID,
		UserID:          rToken.UserID,
		RotationCounter: rToken.RotationCounter,
		IssuedAt:        rToken.IssuedAt.Time,
		ExpiresAt:       rToken.ExpiresAt.Time,
		DeviceInfo:      rToken.DeviceInfo,
		IpAddress:       rToken.IpAddress,
		RevokedAt:       nil,
	}
	return TokensReturn{
		AccessToken:     accessToken,
		RefreshToken:    refreshTokenStr,
		RefreshTokenRaw: rt,
	}, nil
}
