package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type getFieldParam struct {
	Lid       int32
	FieldName string
}

func (s leaderboardService) getField(ctx context.Context, param getFieldParam) (field, error) {
	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.Lid)
	cachedLb, ok := utils.GetCache[leaderboardWithEntry](s.cache, cacheKeyLBFull)
	if ok {
		for _, field := range cachedLb.Fields {
			if param.FieldName == field.Name {
				return field, nil
			}
		}

		return field{}, errNoField
	} else {
		curField, err := s.repo.GetFieldByLID(ctx, database.GetFieldByLIDParams{
			Lid:       param.Lid,
			FieldName: param.FieldName,
		})
		if errors.Is(err, pgx.ErrNoRows) {
			return field{}, errNoField
		}
		if err != nil {
			return field{}, err
		}

		return field{
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
