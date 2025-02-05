package auth

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/utils"
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s AuthService) Refresh(ctx context.Context, param RefreshParam) (TokensReturn, error) {
	claim, err := utils.ValidateRefreshToken(param.RefreshToken, os.Getenv(constants.EnvKeySecret))
	if err != nil {
		return TokensReturn{}, ErrInvalidToken
	}
	rToken, err := s.repo.GetRefreshToken(ctx, claim.TokenID)
	if err == pgx.ErrNoRows {
		return TokensReturn{}, ErrNoTokenExist
	} else if err != nil {
		return TokensReturn{}, err
	}

	if rToken.RotationCounter != claim.RotationCounter {
		return TokensReturn{}, ErrMismatchRotationCounter
	}
	if rToken.RevokedAt.Valid {
		return TokensReturn{}, ErrTokenRevoked
	}

	user, err := s.repo.GetUserByID(ctx, rToken.UserID)
	if err == pgx.ErrNoRows {
		return TokensReturn{}, ErrNoUser
	} else if err != nil {
		return TokensReturn{}, err
	}

	accessTokenStr, err := utils.MakeAccessTokenJWT(user, os.Getenv(constants.EnvKeySecret), constants.AccessTokenDuration)
	if err != nil {
		return TokensReturn{}, err
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
		return TokensReturn{}, err
	}
	rToken, err = s.repo.UpdateRefreshToken(ctx, database.UpdateRefreshTokenParams{
		ExpiresAt:  rToken.ExpiresAt,
		DeviceInfo: rToken.DeviceInfo,
		IpAddress:  rToken.IpAddress,
		ID:         rToken.ID,
	})
	if err != nil {
		return TokensReturn{}, err
	}
	if rToken.RotationCounter != newRotationCounter {
		return TokensReturn{}, ErrMismatchUpdatedRC
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
		AccessToken:     accessTokenStr,
		RefreshToken:    refreshTokenStr,
		RefreshTokenRaw: rt,
	}, nil
}
