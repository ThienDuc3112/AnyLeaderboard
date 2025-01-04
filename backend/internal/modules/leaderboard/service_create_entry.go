package leaderboard

import (
	"anylbapi/internal/database"
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s leaderboardService) createEntry(ctx context.Context, param createEntryParam) (database.LeaderboardEntry, string, error) {
	entry := make(map[string]any)

	if !param.Leaderboard.AllowAnnonymous && param.User == nil {
		return database.LeaderboardEntry{}, "", errNonAnonymousLeaderboard
	} else if param.DisplayName == "" && param.User == nil {
		return database.LeaderboardEntry{}, "", errNoDisplayName
	}

	// Get LB fields
	fields, err := s.repo.GetLeaderboardFieldsByLID(ctx, param.Leaderboard.ID)
	if err != nil {
		return database.LeaderboardEntry{}, "", err
	}

	// Get option field options
	options := map[string][]string{}
	for _, field := range fields {
		if field.FieldValue == database.FieldTypeOPTION {
			optionsForFields, err := s.repo.GetFieldOptions(ctx, database.GetFieldOptionsParams{
				Lid:       param.Leaderboard.ID,
				FieldName: field.FieldName,
			})

			if err != nil {
				return database.LeaderboardEntry{}, "", err
			}
			options[field.FieldName] = optionsForFields
		}
	}

	foundForRankField := false
	var sortedValue float64

	// Processing fields
	for _, field := range fields {
		var input any = param.Entry[field.FieldName]

		switch field.FieldValue {
		case database.FieldTypeDURATION, database.FieldTypeNUMBER:
			val, ok := input.(float64)
			if !ok {
				if !field.Required {
					continue
				}
				return database.LeaderboardEntry{}, field.FieldName, errRequiredFieldNotExist
			}
			entry[field.FieldName] = val

			if foundForRankField && field.ForRank {
				return database.LeaderboardEntry{}, field.FieldName, errConflictForRankField
			}
			if field.ForRank {
				sortedValue = val
				foundForRankField = true
			}

		case database.FieldTypeTIMESTAMP:
			timeStr, ok := input.(string)
			val, err := time.Parse(time.RFC3339, timeStr)
			if !ok || err != nil {
				if !field.Required {
					continue
				}
				return database.LeaderboardEntry{}, field.FieldName, errRequiredFieldNotExist
			}
			entry[field.FieldName] = val.UnixMilli()

			if foundForRankField && field.ForRank {
				return database.LeaderboardEntry{}, field.FieldName, errConflictForRankField
			}
			if field.ForRank {
				sortedValue = float64(val.UnixMilli())
				foundForRankField = true
			}

		case database.FieldTypeTEXT, database.FieldTypeOPTION:
			val, ok := input.(string)
			if !ok {
				if !field.Required {
					continue
				}
				return database.LeaderboardEntry{}, field.FieldName, errRequiredFieldNotExist
			}
			if field.FieldValue == database.FieldTypeOPTION {
				options := options[field.FieldName]
				if len(options) == 0 {
					return database.LeaderboardEntry{}, field.FieldName, errOptionFieldNoOptions
				}
				isAnOption := false
				for _, option := range options {
					if option == val {
						isAnOption = true
						break
					}
				}
				if !isAnOption {
					return database.LeaderboardEntry{}, field.FieldName, errNotAnOption
				}
			}
			entry[field.FieldName] = val

			if field.ForRank {
				return database.LeaderboardEntry{}, field.FieldName, errUnrankableFieldType
			}
		default:
			return database.LeaderboardEntry{}, field.FieldName, errUnrecognizedField
		}
	}

	// Sanity check
	if !foundForRankField {
		return database.LeaderboardEntry{}, "", errNoForRankField
	}

	entryJson, err := json.Marshal(entry)
	if err != nil {
		return database.LeaderboardEntry{}, "", err
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
		return database.LeaderboardEntry{}, "", err
	}

	return e, "", nil
}
