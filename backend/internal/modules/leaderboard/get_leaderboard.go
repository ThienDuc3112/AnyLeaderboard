package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/utils"
	"context"
	"fmt"
)

func (s LeaderboardService) GetLeaderboard(ctx context.Context, id int32) (models.LeaderboardFull, error) {
	// Check cache
	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, id)
	cachedLb, ok := utils.GetCache[models.LeaderboardFull](s.cache, cacheKeyLBFull)
	var res models.LeaderboardFull
	if ok {
		res = cachedLb
		res.Data = make([]models.Entry, 0)

		for i := range res.Fields {
			field := res.Fields[i]
			if field.Type == string(database.FieldTypeOPTION) {
				// Check options cache
				cacheOptionKey := fmt.Sprintf("%s-%s", c.CachePrefixOptions, field.Name)
				cachedOptions, ok := utils.GetCache[[]string](s.cache, cacheOptionKey)
				if ok {
					field.Options = cachedOptions
					res.Fields[i] = field
					continue
				}

				// Get options
				options, err := s.repo.GetFieldOptions(ctx, int32(field.Id))
				if err != nil {
					return models.LeaderboardFull{}, err
				}
				field.Options = options
				s.cache.SetDefault(cacheOptionKey, options)
			}
		}
		return res, nil
	}

	// Get leaderboard
	rows, err := s.repo.GetLeaderboardFull(ctx, id)
	if err != nil {
		return models.LeaderboardFull{}, err
	}
	if len(rows) == 0 {
		return models.LeaderboardFull{}, ErrNoLeaderboard
	}
	lb := rows[0]

	res = models.LeaderboardFull{
		ID:                   int(lb.ID),
		Name:                 lb.Name,
		Description:          lb.Description,
		CoverImageUrl:        lb.CoverImageUrl.String,
		AllowAnonymous:       lb.AllowAnonymous,
		RequiredVerification: lb.RequireVerification,
		UniqueSubmission:     lb.UniqueSubmission,
		CreatedAt:            lb.CreatedAt.Time,
		Descending:           lb.Descending,
		ExternalLink:         make([]models.ExternalLink, 0),
		Fields:               make([]models.Field, 0),
		CreatorId:            int(lb.Creator),
	}

	fieldSet := make(map[int32]bool)
	linkSet := make(map[int]bool)

	for _, row := range rows {
		if row.FieldName.Valid && !fieldSet[row.FieldID.Int32] {
			fieldSet[row.ID] = true
			field := models.Field{
				Id:         int(row.FieldID.Int32),
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
					options, err := s.repo.GetFieldOptions(ctx, int32(field.Id))
					if err != nil {
						return models.LeaderboardFull{}, err
					}
					field.Options = options
					s.cache.SetDefault(cacheOptionKey, options)
				}
			}

			res.Fields = append(res.Fields, field)
		}

		if row.LinkID.Valid && !linkSet[int(row.LinkID.Int32)] {
			linkSet[int(row.LinkID.Int32)] = true
			res.ExternalLink = append(res.ExternalLink, models.ExternalLink{
				Id:           int(row.LinkID.Int32),
				DisplayValue: row.LinkDisplayValue.String,
				URL:          row.LinkUrl.String,
			})
		}
	}

	creatorUsername, err := s.repo.GetUsernameFromId(ctx, lb.Creator)
	if err != nil {
		return models.LeaderboardFull{}, err
	}
	res.Creator = creatorUsername

	// Cache the data
	s.cache.SetDefault(cacheKeyLBFull, res)
	res.Data = make([]models.Entry, 0)
	return res, nil
}
