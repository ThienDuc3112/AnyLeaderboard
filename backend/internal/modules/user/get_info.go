package user

import (
	"anylbapi/internal/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (s UserService) InfoByName(ctx context.Context, Username string) (models.UserPreview, error) {
	user, err := s.repo.GetUserByUsername(ctx, Username)
	if errors.Is(err, pgx.ErrNoRows) {
		return models.UserPreview{}, ErrNoUsers
	} else if err != nil {
		return models.UserPreview{}, err
	}

	return models.UserPreview{
		CreatedAt:   user.CreatedAt.Time,
		Username:    user.Username,
		Description: user.Description,
		DisplayName: user.Username,
	}, nil
}
