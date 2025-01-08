package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"fmt"
)

func (s leaderboardService) getLeaderboard(ctx context.Context, id int32) (leaderboardWithEntry, error) {
	// Check cache
	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, id)
	cachedLb, ok := utils.GetCache[leaderboardWithEntry](s.cache, cacheKeyLBFull)
	var res leaderboardWithEntry
	if ok {
		res = cachedLb
		res.Data = make([]entry, 0)
		return res, nil
	}

	// Get leaderboard
	rows, err := s.repo.GetLeaderboardFull(ctx, id)
	if err != nil {
		return leaderboardWithEntry{}, err
	}
	if len(rows) == 0 {
		return leaderboardWithEntry{}, errNoLeaderboard
	}
	lb := rows[0]

	res = leaderboardWithEntry{
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

	// Cache the data
	s.cache.SetDefault(cacheKeyLBFull, res)
	res.Data = make([]entry, 0)
	return res, nil
}
