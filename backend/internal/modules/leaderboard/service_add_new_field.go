package leaderboard

import (
	"anylbapi/internal/database"
	"context"
	"encoding/json"
	"time"
)

type addFieldParam struct {
	Lid      int32
	NewField struct {
		field
		defaultValue any
	}
}

func (s leaderboardService) addField(ctx context.Context, param addFieldParam) error {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	field := database.CreateLeadeboardFieldParams{
		Lid:        param.Lid,
		FieldName:  param.NewField.Name,
		FieldValue: database.FieldType(param.NewField.Type),
		FieldOrder: int32(param.NewField.FieldOrder),
		ForRank:    param.NewField.ForRank,
		Required:   param.NewField.Required,
		Hidden:     param.NewField.Hidden,
	}

	if field.ForRank {
		return errCannotAddForRank
	}

	var val any
	var ok bool
	if field.Required {
		input := param.NewField.defaultValue
		switch field.FieldValue {
		case database.FieldTypeDURATION, database.FieldTypeNUMBER:
			val, ok = input.(float64)
			if !ok {
				return errConflictType
			}

		case database.FieldTypeTIMESTAMP:
			timeStr, ok := input.(string)
			val, err = time.Parse(time.RFC3339, timeStr)
			if !ok || err != nil {
				return errConflictType
			}

		case database.FieldTypeTEXT, database.FieldTypeOPTION:
			val, ok = input.(string)
			if !ok {
				return errConflictType
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
					return errConflictType
				}
			}
		default:
			return errUnrecognizedField
		}
	}

	err = tx.CreateLeadeboardField(ctx, field)
	if err != nil {
		return err
	}

	if field.FieldValue == database.FieldTypeOPTION {
		options := make([]database.CreateLeadeboardOptionsParams, 0)
		for _, option := range param.NewField.Options {
			options = append(options, database.CreateLeadeboardOptionsParams{
				Lid:       field.Lid,
				FieldName: field.FieldName,
				Option:    option,
			})
		}

		n, err := tx.CreateLeadeboardOptions(ctx, options)
		if err != nil {
			return err
		}
		if int(n) != len(options) {
			return errUnableToInsertAllOptions
		}
	}

	if field.Required {
		jsonVal, err := json.Marshal(val)
		if err != nil {
			return err
		}
		err = tx.UpdateEntryByLeaderboardId(ctx, database.UpdateEntryByLeaderboardIdParams{
			Path:            []string{field.FieldName},
			CreateIfMissing: true,
			LeaderboardID:   field.Lid,
			Value:           jsonVal,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
