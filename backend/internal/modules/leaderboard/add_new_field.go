package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type AddFieldParam struct {
	Lid          int32
	NewField     models.Field
	DefaultValue any
}

func (s LeaderboardService) AddField(ctx context.Context, param AddFieldParam) error {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	field := database.CreateLeaderboardFieldParams{
		Lid:        param.Lid,
		FieldName:  param.NewField.Name,
		FieldValue: database.FieldType(param.NewField.Type),
		FieldOrder: int32(param.NewField.FieldOrder),
		ForRank:    param.NewField.ForRank,
		Required:   param.NewField.Required,
		Hidden:     param.NewField.Hidden,
	}

	if field.ForRank {
		return ErrCannotAddForRank
	}

	var val any
	var ok bool
	if field.Required {
		input := param.DefaultValue
		switch field.FieldValue {
		case database.FieldTypeDURATION, database.FieldTypeNUMBER:
			val, ok = input.(float64)
			if !ok {
				return ErrConflictType
			}

		case database.FieldTypeTIMESTAMP:
			timeStr, ok := input.(string)
			val, err = time.Parse(time.RFC3339, timeStr)
			if !ok || err != nil {
				return ErrConflictType
			}

		case database.FieldTypeTEXT, database.FieldTypeOPTION:
			val, ok = input.(string)
			if !ok {
				return ErrConflictType
			}
			if field.FieldValue == database.FieldTypeOPTION {
				matched := false
				for _, option := range param.NewField.Options {
					if option == val {
						matched = true
						break
					}
				}

				if !matched {
					return ErrConflictType
				}
			}
		default:
			return ErrUnrecognizedField
		}
	}

	fid, err := tx.CreateLeaderboardField(ctx, field)

	if err != nil {
		return err
	}

	if field.FieldValue == database.FieldTypeOPTION {
		options := make([]database.CreateLeaderboardOptionsParams, 0)
		for _, option := range param.NewField.Options {
			options = append(options, database.CreateLeaderboardOptionsParams{
				Fid:    fid,
				Option: option,
			})
		}

		n, err := tx.CreateLeaderboardOptions(ctx, options)
		if err != nil {
			return err
		}
		if int(n) != len(options) {
			return ErrUnableToInsertAllOptions
		}
	}

	if field.Required {
		jsonVal, err := json.Marshal(val)
		if err != nil {
			return err
		}
		err = tx.AddFieldToEntriesByLeaderboardId(ctx, database.AddFieldToEntriesByLeaderboardIdParams{
			Path:            []string{field.FieldName},
			CreateIfMissing: true,
			LeaderboardID:   field.Lid,
			Value:           jsonVal,
		})
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	cacheKeyLeaderboard := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.Lid)
	s.cache.Delete(cacheKeyLeaderboard)

	return nil
}
