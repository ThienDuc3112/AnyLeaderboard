package auth

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"database/sql"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s authService) login(context context.Context, body loginParam) (loginReturn, error) {
	loginWithEmail := false
	if strings.Contains(body.Username, "@") {
		loginWithEmail = true
	}

	var user database.User
	var err error
	if loginWithEmail {
		user, err = s.repo.GetUserByEmail(context, body.Username)
	} else {
		user, err = s.repo.GetUserByEmail(context, body.Username)
	}

	if err == sql.ErrNoRows {
		return loginReturn{}, errNoUser
	} else if err != nil {
		return loginReturn{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return loginReturn{}, errIncorrectPassword
	}

	accessToken, err := utils.MakeJWT(user, os.Getenv(constants.EnvKeySecret), time.Minute*30)
	if err != nil {
		return loginReturn{}, err
	}

	refreshTokenParam := database.CreateNewRefreshTokenParams{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(14 * 24 * time.Hour),
	}
	if body.DeviceInfo != "" {
		refreshTokenParam.DeviceInfo = sql.NullString{
			String: body.DeviceInfo,
			Valid:  true,
		}
	}
	if body.IpAddress != "" {
		refreshTokenParam.IpAddress = sql.NullString{
			String: body.IpAddress,
			Valid:  true,
		}
	}

	refreshToken, err := s.repo.CreateNewRefreshToken(context, refreshTokenParam)
	if err != nil {
		return loginReturn{}, err
	}

	refreshTokenStr, err := utils.MakeRefreshTokenJWT(refreshToken, os.Getenv(constants.EnvKeySecret), refreshToken.ExpiresAt)
	if err != nil {
		return loginReturn{}, err
	}

	return loginReturn{
		accessToken:     accessToken,
		refreshToken:    refreshTokenStr,
		refreshTokenRaw: refreshToken,
	}, nil
}
