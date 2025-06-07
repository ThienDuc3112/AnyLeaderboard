package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"slices"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s LeaderboardService) CreateEntry(ctx context.Context, param CreateEntryParam) (database.LeaderboardEntry, string, error) {
	entryData := make(map[string]any)

	if !param.Leaderboard.AllowAnonymous && param.User == nil {
		return database.LeaderboardEntry{}, "", ErrNonAnonymousLeaderboard
	} else if param.DisplayName == "" && param.User == nil {
		return database.LeaderboardEntry{}, "", ErrNoDisplayName
	}

	// Get LB fields
	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.Leaderboard.ID)
	cachedLb, ok := utils.GetCache[models.LeaderboardFull](s.cache, cacheKeyLBFull)
	var lb models.LeaderboardFull
	if ok {
		lb = cachedLb
		lb.Data = make([]models.Entry, 0)
	} else {
		var err error
		lb, err = s.GetLeaderboard(ctx, int32(param.Leaderboard.ID))
		if err != nil {
			return database.LeaderboardEntry{}, "", err
		}

		s.cache.SetDefault(cacheKeyLBFull, lb)
		lb.Data = make([]models.Entry, 0)
	}

	fmt.Printf("Leaderboard: %+v\n", lb)

	foundForRankField := false
	var sortedValue float64

	// Processing fields
	fmt.Println()
	for _, field := range lb.Fields {
		var input any = param.Entry[field.Name]

		fmt.Printf("body: %+v\tfield name: %+v\tinput: %+v\n", param.Entry, field.Name, input)
		switch database.FieldType(field.Type) {
		case database.FieldTypeDURATION, database.FieldTypeNUMBER, database.FieldTypeTIMESTAMP:
			val, ok := input.(float64)
			if !ok {
				if !field.Required {
					continue
				}
				return database.LeaderboardEntry{}, field.Name, ErrRequiredFieldNotExist
			}
			entryData[fmt.Sprintf("%d", field.Id)] = val

			if foundForRankField && field.ForRank {
				return database.LeaderboardEntry{}, field.Name, ErrConflictForRankField
			}
			if field.ForRank {
				sortedValue = val
				foundForRankField = true
			}

		case database.FieldTypeTEXT, database.FieldTypeOPTION:
			val, ok := input.(string)
			if !ok {
				if !field.Required {
					continue
				}
				return database.LeaderboardEntry{}, field.Name, ErrRequiredFieldNotExist
			}
			if database.FieldType(field.Type) == database.FieldTypeOPTION {
				options := field.Options
				if len(options) == 0 {
					return database.LeaderboardEntry{}, field.Name, ErrOptionFieldNoOptions
				}

				isAnOption := slices.Contains(options, val)

				if !isAnOption {
					return database.LeaderboardEntry{}, field.Name, ErrNotAnOption
				}
			}
			entryData[fmt.Sprintf("%d", field.Id)] = val

			if field.ForRank {
				return database.LeaderboardEntry{}, field.Name, ErrUnrankableFieldType
			}
		default:
			return database.LeaderboardEntry{}, field.Name, ErrUnrecognizedField
		}
	}

	// Sanity check
	if !foundForRankField {
		return database.LeaderboardEntry{}, "", ErrNoForRankField
	}

	entryJson, err := json.Marshal(entryData)
	if err != nil {
		return database.LeaderboardEntry{}, "", err
	}

	userId := pgtype.Int4{}
	if param.DisplayName == "" {
		userId.Int32 = param.User.ID
		userId.Valid = true
	}

	entryParam := database.CreateLeaderboardEntryParams{
		CustomFields:  entryJson,
		UserID:        userId,
		Username:      param.DisplayName,
		LeaderboardID: int32(param.Leaderboard.ID),
		SortedField:   sortedValue,
	}

	e, err := s.repo.CreateLeaderboardEntry(ctx, entryParam)
	if err != nil {
		return database.LeaderboardEntry{}, "", err
	}

	return e, "", nil
}
