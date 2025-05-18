package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/utils"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s LeaderboardService) CreateLeaderboard(ctx context.Context, param CreateLeaderboardParam) (models.Leaderboard, error) {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return models.Leaderboard{}, err
	}
	defer tx.Rollback(ctx)

	// ================== Processing leaderboard ==================
	lbParam := database.CreateLeaderboardParams{
		Name:                param.Name,
		Creator:             param.User.ID,
		Description:         param.Description,
		AllowAnonymous:      param.AllowAnonymous,
		RequireVerification: param.RequiredVerification,
		UniqueSubmission:    param.UniqueSubmission,
		Descending:          param.Descending,
		NameLanguage:        utils.DetectLanguage(param.Name),
		DescriptionLanguage: utils.DetectLanguage(param.Description),
	}
	if param.CoverImageURL != "" {
		lbParam.CoverImageUrl = pgtype.Text{
			String: param.CoverImageURL,
			Valid:  true,
		}
	}

	// ================== Inserting leaderboard ==================
	lb, err := tx.CreateLeaderboard(ctx, lbParam)
	if err != nil {
		return models.Leaderboard{}, err
	}

	// ================== Processing leaderboard links ==================
	links := make([]database.CreateLeaderboardExternalLinkParams, 0)
	for _, link := range param.ExternalLinks {
		links = append(links, database.CreateLeaderboardExternalLinkParams{
			LeaderboardID: lb.ID,
			DisplayValue:  link.DisplayValue,
			Url:           link.URL,
		})
	}

	// ================== Inserting leaderboard links ==================
	if len(links) > 0 {
		n, err := tx.CreateLeaderboardExternalLink(ctx, links)
		if err != nil {
			return models.Leaderboard{}, err
		}
		if n != int64(len(links)) {
			return models.Leaderboard{}, ErrUnableToInsertAllLinks
		}
	}

	// ================== Processing leaderboard fields ==================
	type optionsToInsert struct {
		fieldName string
		lid       int32
		option    []string
	}
	fields := make([]database.CreateLeadeboardFieldsParams, 0)
	opts := make([]optionsToInsert, 0)
	forRankExist := false
	nonHiddenFieldExist := false

	for _, field := range param.Fields {
		if forRankExist && field.ForRank {
			return models.Leaderboard{}, ErrMultipleForRankField
		}
		if field.ForRank && !field.Required {
			return models.Leaderboard{}, ErrForRankNotRequired
		}
		forRankExist = forRankExist || field.ForRank
		nonHiddenFieldExist = nonHiddenFieldExist || !field.Hidden
		fields = append(fields, database.CreateLeadeboardFieldsParams{
			Lid:        lb.ID,
			FieldName:  field.Name,
			FieldValue: database.FieldType(field.Type),
			FieldOrder: int32(field.FieldOrder),
			ForRank:    field.ForRank,
			Required:   field.Required,
			Hidden:     field.Hidden,
		})

		if field.Type == string(database.FieldTypeOPTION) {
			if len(field.Options) == 0 {
				return models.Leaderboard{}, ErrNoOptions
			}
			opts = append(opts, optionsToInsert{
				fieldName: field.Name,
				lid:       lb.ID,
				option:    field.Options,
			})
		}
	}

	if !forRankExist {
		return models.Leaderboard{}, ErrNoForRankField
	}
	if !nonHiddenFieldExist {
		return models.Leaderboard{}, ErrNoPublicField
	}

	// ================== Inserting leaderboard fields ==================
	n, err := tx.CreateLeadeboardFields(ctx, fields)
	if err != nil {
		return models.Leaderboard{}, err
	}
	if n != int64(len(fields)) {
		return models.Leaderboard{}, ErrUnableToInsertAllFields
	}

	for _, opt := range opts {
		// ================== Processing leaderboard options ==================
		optParam := make([]database.CreateLeadeboardOptionsParams, 0)
		for _, o := range opt.option {
			optParam = append(optParam, database.CreateLeadeboardOptionsParams{
				Lid:       lb.ID,
				FieldName: opt.fieldName,
				Option:    o,
			})
		}

		// ================== Inserting leaderboard options ==================
		n, err = tx.CreateLeadeboardOptions(ctx, optParam)
		if err != nil {
			return models.Leaderboard{}, err
		}
		if n != int64(len(optParam)) {
			return models.Leaderboard{}, ErrUnableToInsertAllOptions
		}
	}

	// ================== Commiting changes ==================
	err = tx.Commit(ctx)
	if err != nil {
		return models.Leaderboard{}, err
	}

	// Remove caching if exist
	s.cache.Delete(fmt.Sprintf("%s-%d", c.CachePrefixNoLeaderboard, lb.ID))
	return models.Leaderboard{
		ID:                   int(lb.ID),
		Name:                 lb.Name,
		Description:          lb.Description,
		CoverImageUrl:        lb.CoverImageUrl.String,
		CreatedAt:            lb.CreatedAt.Time,
		Creator:              int(lb.Creator),
		UpdatedAt:            lb.UpdatedAt.Time,
		AllowAnonymous:       lb.AllowAnonymous,
		RequiredVerification: lb.RequireVerification,
		UniqueSubmission:     lb.UniqueSubmission,
	}, nil
}
