package auth

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func (s authService) login(context context.Context, param loginParam) (loginReturn, error) {
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
		return loginReturn{}, errNoUser
	} else if err != nil {
		return loginReturn{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))
	if err != nil {
		return loginReturn{}, errIncorrectPassword
	}

	accessToken, err := utils.MakeJWT(user, os.Getenv(constants.EnvKeySecret), time.Minute*30)
	if err != nil {
		return loginReturn{}, err
	}

	refreshTokenParam := database.CreateNewRefreshTokenParams{
		UserID: user.ID,
		ExpiresAt: pgtype.Timestamp{
			Time:  time.Now().Add(14 * 24 * time.Hour),
			Valid: true,
		},
		DeviceInfo: param.DeviceInfo,
		IpAddress:  param.IpAddress,
	}

	refreshToken, err := s.repo.CreateNewRefreshToken(context, refreshTokenParam)
	if err != nil {
		return loginReturn{}, err
	}

	refreshTokenStr, err := utils.MakeRefreshTokenJWT(refreshToken, os.Getenv(constants.EnvKeySecret), refreshToken.ExpiresAt.Time)
	if err != nil {
		return loginReturn{}, err
	}

	return loginReturn{
		accessToken:     accessToken,
		refreshToken:    refreshTokenStr,
		refreshTokenRaw: refreshToken,
	}, nil
}
