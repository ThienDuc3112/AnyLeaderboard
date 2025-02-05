package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/utils"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type GetFieldParam struct {
	Lid       int32
	FieldName string
}

func (s LeaderboardService) GetField(ctx context.Context, param GetFieldParam) (models.Field, error) {
	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.Lid)
	cachedLb, ok := utils.GetCache[models.LeaderboardFull](s.cache, cacheKeyLBFull)
	if ok {
		for _, field := range cachedLb.Fields {
			if param.FieldName == field.Name {
				return field, nil
			}
		}

		return models.Field{}, ErrNoField
	} else {
		curField, err := s.repo.GetFieldByLID(ctx, database.GetFieldByLIDParams{
			Lid:       param.Lid,
			FieldName: param.FieldName,
		})
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Field{}, ErrNoField
		}
		if err != nil {
			return models.Field{}, err
		}

		return models.Field{
			Name:       curField.FieldName,
			Required:   curField.Required,
			Hidden:     curField.Hidden,
			FieldOrder: int(curField.FieldOrder),
			Type:       string(curField.FieldValue),
			Options:    nil,
			ForRank:    curField.ForRank,
		}, nil
	}
}
