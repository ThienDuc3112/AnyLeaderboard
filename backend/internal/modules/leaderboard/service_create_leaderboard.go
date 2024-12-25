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
	lb, err := tx.CreateLeaderboard(ctx, lbParam)
	if err != nil {
		return database.Leaderboard{}, err
	}

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
		if field.Type == "OPTION" {
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

	n, err := tx.CreateLeadeboardFields(ctx, fields)
	if err != nil {
		return database.Leaderboard{}, err
	}
	if n != int64(len(fields)) {
		return database.Leaderboard{}, errUnableToInsertAllFields
	}

	for _, opt := range opts {
		optParam := make([]database.CreateLeadeboardOptionsParams, 0)
		for _, o := range opt.option {
			optParam = append(optParam, database.CreateLeadeboardOptionsParams{
				Lid:       lb.ID,
				FieldName: opt.fieldName,
				Option:    o,
			})
		}
		n, err = tx.CreateLeadeboardOptions(ctx, optParam)
		if err != nil {
			return database.Leaderboard{}, err
		}
		if n != int64(len(optParam)) {
			return database.Leaderboard{}, errUnableToInsertAllOptions
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return database.Leaderboard{}, err
	}

	return lb, err
}
