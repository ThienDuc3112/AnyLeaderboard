package favorite

import (
	"anylbapi/internal/database"
	"context"
)

type DeleteParam struct {
	Uid int32
	Lid int32
}

func (s *FavoriteService) DeleteFavorite(ctx context.Context, param DeleteParam) error {
	return s.repo.DeleteFavorite(ctx, database.DeleteFavoriteParams{
		UserID:        param.Uid,
		LeaderboardID: param.Lid,
	})
}
