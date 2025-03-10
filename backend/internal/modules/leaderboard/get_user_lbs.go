package leaderboard

import (
	"anylbapi/internal/models"
	"context"
	"fmt"
)

type GetByUserParam struct {
	UserId int32
}

func (s LeaderboardService) GetByUser(ctx context.Context, param GetByUserParam) ([]models.LeaderboardPreview, error) {
	return nil, fmt.Errorf("Unimplemented")
}
