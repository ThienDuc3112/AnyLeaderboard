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
	fields := database.BulkInsertFieldsParams{
		Lids:        make([]int32, 0),
		FieldNames:  make([]string, 0),
		FieldValues: make([]string, 0),
		FieldOrders: make([]int32, 0),
		ForRanks:    make([]bool, 0),
		Required:    make([]bool, 0),
		Hidden:      make([]bool, 0),
	}
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
		fields.Lids = append(fields.Lids, lb.ID)
		fields.FieldNames = append(fields.FieldNames, field.Name)
		fields.FieldValues = append(fields.FieldValues, field.Type)
		fields.FieldOrders = append(fields.FieldOrders, int32(field.FieldOrder))
		fields.ForRanks = append(fields.ForRanks, field.ForRank)
		fields.Required = append(fields.Required, field.Required)
		fields.Hidden = append(fields.Hidden, field.Hidden)
	}

	if !forRankExist {
		return models.Leaderboard{}, ErrNoForRankField
	}
	if !nonHiddenFieldExist {
		return models.Leaderboard{}, ErrNoPublicField
	}

	// ================== Inserting leaderboard fields ==================
	fieldIds, err := tx.BulkInsertFields(ctx, fields)
	if err != nil {
		return models.Leaderboard{}, err
	}
	if len(fieldIds) != len(param.Fields) {
		return models.Leaderboard{}, ErrUnableToInsertAllFields
	}

	for i, field := range param.Fields {
		// ================== Processing leaderboard options ==================
		if field.Type != string(database.FieldTypeOPTION) {
			continue
		}
		if len(field.Options) == 0 {
			return models.Leaderboard{}, ErrNoOptions
		}
		optParam := make([]database.CreateLeadeboardOptionsParams, 0)
		for _, o := range field.Options {
			optParam = append(optParam, database.CreateLeadeboardOptionsParams{
				Fid:    fieldIds[i],
				Option: o,
			})
		}

		// ================== Inserting leaderboard options ==================
		n, err := tx.CreateLeadeboardOptions(ctx, optParam)
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
