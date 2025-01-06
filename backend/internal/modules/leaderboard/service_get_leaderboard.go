package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"fmt"
)

func (s leaderboardService) getLeaderboard(ctx context.Context, param getLeaderboardParam) (leaderboardWithEntry, error) {
	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.id)
	cachedLb, ok := utils.GetCache[leaderboardWithEntry](s.cache, cacheKeyLBFull)

	var res leaderboardWithEntry
	if ok {
		res = cachedLb
		res.Data = make([]entry, 0)
	} else {
		var err error
		res, err = s.__getLeaderboard(ctx, int32(param.id))
		if err != nil {
			return leaderboardWithEntry{}, err
		}

		s.cache.SetDefault(cacheKeyLBFull, res)
		res.Data = make([]entry, 0)
	}

	var entries []database.LeaderboardEntry
	var count int64
	var err error
	if res.RequiredVerification {
		entries, err = s.repo.GetVerifiedEntriesFromLeaderboardId(ctx, database.GetVerifiedEntriesFromLeaderboardIdParams{
			LeaderboardID: int32(res.ID),
			Offset:        int32(param.offset),
			Limit:         int32(param.pageSize),
		})
		if err != nil {
			return leaderboardWithEntry{}, err
		}

		count, err = s.repo.GetLeaderboardVerifiedEntriesCount(ctx, int32(res.ID))
		if err != nil {
			return leaderboardWithEntry{}, err
		}
	} else {
		entries, err = s.repo.GetEntriesFromLeaderboardId(ctx, database.GetEntriesFromLeaderboardIdParams{
			LeaderboardID: int32(res.ID),
			Offset:        int32(param.offset),
			Limit:         int32(param.pageSize),
		})
		if err != nil {
			return leaderboardWithEntry{}, err
		}

		count, err = s.repo.GetLeaderboardEntriesCount(ctx, int32(res.ID))
		if err != nil {
			return leaderboardWithEntry{}, err
		}
	}

	res.EntriesCount = int(count)

	for _, row := range entries {
		entry := entry{
			Id:        int(row.ID),
			CreatedAt: row.CreatedAt.Time,
			UpdatedAt: row.UpdatedAt.Time,
			Fields:    row.CustomFields,
		}

		res.Data = append(res.Data, entry)
	}

	return res, nil
}

func (s leaderboardService) __getLeaderboard(ctx context.Context, id int32) (leaderboardWithEntry, error) {
	rows, err := s.repo.GetLeaderboardFull(ctx, id)
	if err != nil {
		return leaderboardWithEntry{}, err
	}
	if len(rows) == 0 {
		return leaderboardWithEntry{}, errNoLeaderboard
	}
	lb := rows[0]

	res := leaderboardWithEntry{
		ID:                   int(lb.ID),
		Name:                 lb.Name,
		Description:          lb.Description,
		CoverImageUrl:        lb.CoverImageUrl.String,
		AllowAnonymous:       lb.AllowAnnonymous,
		RequiredVerification: lb.RequireVerification,
		UniqueSubmission:     lb.UniqueSubmission,
		ExternalLink:         make([]externalLink, 0),
		Fields:               make([]field, 0),
	}

	fieldSet := make(map[string]bool)
	linkSet := make(map[int]bool)

	for _, row := range rows {
		if row.FieldName.Valid && !fieldSet[row.FieldName.String] {
			fieldSet[row.FieldName.String] = true
			field := field{
				Name:       row.FieldName.String,
				Type:       string(row.FieldValue.FieldType),
				Required:   row.FieldRequired.Bool,
				Hidden:     row.FieldHidden.Bool,
				FieldOrder: int(row.FieldOrder.Int32),
				ForRank:    row.FieldForRank.Bool,
			}

			if field.Type == string(database.FieldTypeOPTION) {
				cacheOptionKey := fmt.Sprintf("%s-%s", c.CachePrefixOptions, field.Name)
				cached := false
				if cachedOptions, exist := s.cache.Get(cacheOptionKey); exist {
					if options, ok := cachedOptions.([]string); ok {
						field.Options = options
						cached = true
					} else {
						s.cache.Delete(cacheOptionKey)
					}
				}

				if !cached {
					options, err := s.repo.GetFieldOptions(ctx, database.GetFieldOptionsParams{
						Lid:       lb.ID,
						FieldName: field.Name,
					})
					if err != nil {
						return leaderboardWithEntry{}, err
					}
					field.Options = options
					s.cache.SetDefault(cacheOptionKey, options)
				}
			}

			res.Fields = append(res.Fields, field)
		}

		if row.LinkID.Valid && !linkSet[int(row.LinkID.Int32)] {
			linkSet[int(row.LinkID.Int32)] = true
			res.ExternalLink = append(res.ExternalLink, externalLink{
				DisplayValue: row.LinkDisplayValue.String,
				URL:          row.LinkUrl.String,
			})
		}
	}
	return res, nil
}
