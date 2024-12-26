package leaderboard

import (
	"anylbapi/internal/database"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s leaderboardService) createLeaderboard(ctx context.Context, param createLeaderboardParam) (database.Leaderboard, error) {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return database.Leaderboard{}, nil
	}
	defer tx.Rollback(ctx)

	// ================== Processing leaderboard ==================
	lbParam := database.CreateLeaderboardParams{
		Name:                param.Name,
		Creator:             param.User.ID,
		Description:         param.Description,
		AllowAnnonymous:     param.AllowAnonymous,
		RequireVerification: param.RequiredVerification,
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
		return database.Leaderboard{}, err
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
			return database.Leaderboard{}, err
		}
		if n != int64(len(links)) {
			return database.Leaderboard{}, errUnableToInsertAllLinks
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
			return database.Leaderboard{}, errMultipleForRankField
		}
		forRankExist = forRankExist || field.ForRank
		nonHiddenFieldExist = nonHiddenFieldExist || !field.Hidden
		fields = append(fields, database.CreateLeadeboardFieldsParams{
			Lid:        lb.ID,
			FieldName:  field.Name,
			FieldValue: database.FieldType(field.Type),
			FieldOrder: int32(field.FieldOrder),
			ForRank:    field.ForRank,
		})
		if field.Type == string(database.FieldTypeOPTION) {
			if len(field.Options) == 0 {
				return database.Leaderboard{}, errNoOptions
			}
			opts = append(opts, optionsToInsert{
				fieldName: field.Name,
				lid:       lb.ID,
				option:    field.Options,
			})
		}
	}
	if !forRankExist {
		return database.Leaderboard{}, errNoForRankField
	}
	if !nonHiddenFieldExist {
		return database.Leaderboard{}, errNoPublicField
	}

	// ================== Inserting leaderboard fields ==================
	n, err := tx.CreateLeadeboardFields(ctx, fields)
	if err != nil {
		return database.Leaderboard{}, err
	}
	if n != int64(len(fields)) {
		return database.Leaderboard{}, errUnableToInsertAllFields
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
			return database.Leaderboard{}, err
		}
		if n != int64(len(optParam)) {
			return database.Leaderboard{}, errUnableToInsertAllOptions
		}
	}

	// ================== Commiting changes ==================
	err = tx.Commit(ctx)
	if err != nil {
		return database.Leaderboard{}, err
	}

	return lb, err
}
