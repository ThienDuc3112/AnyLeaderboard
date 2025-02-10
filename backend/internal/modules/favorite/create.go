package favorite

import (
	"anylbapi/internal/database"
	"context"
)

type CreateParam struct {
	Uid int32
	Lid int32
}

func (s *FavoriteService) Create(ctx context.Context, param CreateParam) error {
	return s.repo.AddFavorite(ctx, database.AddFavoriteParams{
		UserID:        param.Uid,
		LeaderboardID: param.Lid,
	})
}
