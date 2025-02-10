package favorite

import (
	"context"
)

type GetByUserParam struct {
	Uid int32
}

func (s *FavoriteService) GetByUser(ctx context.Context, param GetByUserParam) ([]int32, error) {
	favs, err := s.repo.GetUserFavorite(ctx, param.Uid)
	if err != nil {
		return nil, err
	}

	res := make([]int32, len(favs))
	for i, fav := range favs {
		res[i] = fav.LeaderboardID
	}

	return res, nil
}
