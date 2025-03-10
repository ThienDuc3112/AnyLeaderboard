package user

import (
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type UpdateParam struct {
	Username    string
	Description string
	DisplayName string
}

func (s UserService) Update(ctx context.Context, param UpdateParam) (models.UserPreview, error) {
	user, err := s.repo.UpdateUser(ctx, database.UpdateUserParams{
		Description: param.Description,
		DisplayName: param.DisplayName,
		Username:    param.Username,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return models.UserPreview{}, ErrNoUsers
	} else if err != nil {
		return models.UserPreview{}, err
	}

	return models.UserPreview{
		CreatedAt:   user.CreatedAt.Time,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Description: user.Description,
	}, nil

}
