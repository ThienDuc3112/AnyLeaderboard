package leaderboard

import (
	"anylbapi/internal/database"
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s leaderboardService) createEntry(ctx context.Context, param createEntryParam) (database.LeaderboardEntry, error, string) {
	entry := make(map[string]any)

	if !param.Leaderboard.AllowAnnonymous && param.User == nil {
		return database.LeaderboardEntry{}, errNonAnonymousLeaderboard, ""
	} else if param.DisplayName == "" && param.User == nil {
		return database.LeaderboardEntry{}, errNonAnonymousLeaderboard, ""
	}

	foundForRankField := false
	var sortedValue float64

	for _, field := range param.Fields {
		var input any = param.Entry[field.FieldName]

		switch field.FieldValue {
		case database.FieldTypeDURATION, database.FieldTypeNUMBER:
			val, ok := input.(float64)
			if !ok {
				if !field.Required {
					continue
				}
				return database.LeaderboardEntry{}, errRequiredFieldNotExist, field.FieldName
			}
			entry[field.FieldName] = val

			if foundForRankField && field.ForRank {
				return database.LeaderboardEntry{}, errConflictForRankField, field.FieldName
			}
			if field.ForRank {
				sortedValue = val
				foundForRankField = true
			}

		case database.FieldTypeTIMESTAMP:
			val, ok := input.(time.Time)
			if !ok {
				if !field.Required {
					continue
				}
				return database.LeaderboardEntry{}, errRequiredFieldNotExist, field.FieldName
			}
			entry[field.FieldName] = val.Unix()

			if foundForRankField && field.ForRank {
				return database.LeaderboardEntry{}, errConflictForRankField, field.FieldName
			}
			if field.ForRank {
				sortedValue = float64(val.Unix())
				foundForRankField = true
			}

		case database.FieldTypeTEXT, database.FieldTypeOPTION:
			val, ok := input.(string)
			if !ok {
				if !field.Required {
					continue
				}
				return database.LeaderboardEntry{}, errRequiredFieldNotExist, field.FieldName
			}
			if field.FieldValue == database.FieldTypeOPTION {
				options := param.Options[field.FieldName]
				if len(options) == 0 {
					return database.LeaderboardEntry{}, errOptionFieldNoOptions, field.FieldName
				}
				isAnOption := false
				for _, option := range options {
					if option.Option == val {
						isAnOption = true
						break
					}
				}
				if !isAnOption {
					return database.LeaderboardEntry{}, errNotAnOption, field.FieldName
				}
			}
			entry[field.FieldName] = val

			if field.ForRank {
				return database.LeaderboardEntry{}, errUnrankableFieldType, field.FieldName
			}
		default:
			return database.LeaderboardEntry{}, errUnrecognizedField, field.FieldName
		}
	}

	if !foundForRankField {
		return database.LeaderboardEntry{}, errNoForRankField, ""
	}

	entryJson, err := json.Marshal(entry)
	if err != nil {
		return database.LeaderboardEntry{}, err, ""
	}

	userId := pgtype.Int4{}
	if param.DisplayName == "" {
		userId.Int32 = param.User.ID
		userId.Valid = true
	}

	entryParam := database.CreateLeadeboardEntryParams{
		CustomFields:  entryJson,
		UserID:        userId,
		Username:      param.DisplayName,
		LeaderboardID: param.Leaderboard.ID,
		SortedField:   sortedValue,
	}

	e, err := s.repo.CreateLeadeboardEntry(ctx, entryParam)
	if err != nil {
		return database.LeaderboardEntry{}, err, ""
	}

	return e, nil, ""
}
