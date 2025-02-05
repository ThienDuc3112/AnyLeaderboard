package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s leaderboardService) createEntry(ctx context.Context, param createEntryParam) (database.LeaderboardEntry, string, error) {
	entryData := make(map[string]any)

	if !param.Leaderboard.AllowAnnonymous && param.User == nil {
		return database.LeaderboardEntry{}, "", ErrNonAnonymousLeaderboard
	} else if param.DisplayName == "" && param.User == nil {
		return database.LeaderboardEntry{}, "", ErrNoDisplayName
	}

	// Get LB fields
	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.Leaderboard.ID)
	cachedLb, ok := utils.GetCache[leaderboardWithEntry](s.cache, cacheKeyLBFull)
	var lb leaderboardWithEntry
	if ok {
		lb = cachedLb
		lb.Data = make([]entry, 0)
	} else {
		var err error
		lb, err = s.getLeaderboard(ctx, param.Leaderboard.ID)
		if err != nil {
			return database.LeaderboardEntry{}, "", err
		}

		s.cache.SetDefault(cacheKeyLBFull, lb)
		lb.Data = make([]entry, 0)
	}

	foundForRankField := false
	var sortedValue float64

	// Processing fields
	for _, field := range lb.Fields {
		var input any = param.Entry[field.Name]

		switch database.FieldType(field.Type) {
		case database.FieldTypeDURATION, database.FieldTypeNUMBER:
			val, ok := input.(float64)
			if !ok {
				if !field.Required {
					continue
				}
				return database.LeaderboardEntry{}, field.Name, ErrRequiredFieldNotExist
			}
			entryData[field.Name] = val

			if foundForRankField && field.ForRank {
				return database.LeaderboardEntry{}, field.Name, ErrConflictForRankField
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
				return database.LeaderboardEntry{}, field.Name, ErrRequiredFieldNotExist
			}
			entryData[field.Name] = val.UnixMilli()

			if foundForRankField && field.ForRank {
				return database.LeaderboardEntry{}, field.Name, ErrConflictForRankField
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
				return database.LeaderboardEntry{}, field.Name, ErrRequiredFieldNotExist
			}
			if database.FieldType(field.Type) == database.FieldTypeOPTION {
				options := field.Options
				if len(options) == 0 {
					return database.LeaderboardEntry{}, field.Name, ErrOptionFieldNoOptions
				}
				isAnOption := false
				for _, option := range options {
					if option == val {
						isAnOption = true
						break
					}
				}
				if !isAnOption {
					return database.LeaderboardEntry{}, field.Name, ErrNotAnOption
				}
			}
			entryData[field.Name] = val

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
