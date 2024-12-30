package auth

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s authService) refresh(ctx context.Context, param refreshParam) (tokensReturn, error) {
	claim, err := utils.ValidateRefreshToken(param.RefreshToken, os.Getenv(constants.EnvKeySecret))
	if err != nil {
		return tokensReturn{}, errInvalidToken
	}
	rToken, err := s.repo.GetRefreshToken(ctx, claim.TokenID)
	if err == pgx.ErrNoRows {
		return tokensReturn{}, errNoTokenExist
	} else if err != nil {
		return tokensReturn{}, err
	}

	if rToken.RotationCounter != claim.RotationCounter {
		return tokensReturn{}, errMismatchRotationCounter
	}
	if rToken.RevokedAt.Valid {
		return tokensReturn{}, errTokenRevoked
	}

	user, err := s.repo.GetUserByID(ctx, rToken.UserID)
	if err == pgx.ErrNoRows {
		return tokensReturn{}, errNoUser
	} else if err != nil {
		return tokensReturn{}, err
	}

	accessTokenStr, err := utils.MakeAccessTokenJWT(user, os.Getenv(constants.EnvKeySecret), constants.AccessTokenDuration)
	if err != nil {
		return tokensReturn{}, err
	}
	rToken.ExpiresAt = pgtype.Timestamptz{
		Time:  time.Now().Add(constants.RefreshTokenDuration),
		Valid: true,
	}
	newRotationCounter := rToken.RotationCounter + 1
	rToken.RotationCounter = newRotationCounter
	rToken.DeviceInfo = param.DeviceInfo
	rToken.IpAddress = param.IpAddress
	refreshTokenStr, err := utils.MakeRefreshTokenJWT(rToken, os.Getenv(constants.EnvKeySecret), rToken.ExpiresAt.Time)
	if err != nil {
		return tokensReturn{}, err
	}
	rToken, err = s.repo.UpdateRefreshToken(ctx, database.UpdateRefreshTokenParams{
		ExpiresAt:  rToken.ExpiresAt,
		DeviceInfo: rToken.DeviceInfo,
		IpAddress:  rToken.IpAddress,
		ID:         rToken.ID,
	})
	if err != nil {
		return tokensReturn{}, err
	}
	if rToken.RotationCounter != newRotationCounter {
		return tokensReturn{}, errMismatchUpdatedRC
	}

	return tokensReturn{
		accessToken:     accessTokenStr,
		refreshToken:    refreshTokenStr,
		refreshTokenRaw: rToken,
	}, nil
}
