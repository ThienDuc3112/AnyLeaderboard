package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gookit/goutil/arrutil"
	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateLeaderboardParam struct {
	models.LeaderboardFull
	UserId int
}

var (
	ErrUnknownEditField  = fmt.Errorf("try to edit a non-existent field")
	ErrChangeFieldType   = fmt.Errorf("changing field type is forbidden")
	ErrChangeForRank     = fmt.Errorf("changing for rank field is forbidden")
	ErrHiddenForRank     = fmt.Errorf("hidden for rank field is forbidden")
	ErrNoDefault         = fmt.Errorf("no default for new required were set")
	ErrNotOwnLeaderboard = fmt.Errorf("cannot delete other's leaderboard")
)

func (s LeaderboardService) UpdateLeaderboard(ctx context.Context, param UpdateLeaderboardParam) error {
	oldLb, err := s.GetLeaderboard(ctx, int32(param.ID))
	if err != nil {
		return err
	}

	if param.UserId != oldLb.CreatorId {
		return ErrNotOwnLeaderboard
	}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = tx.UpdateLeaderboard(ctx, database.UpdateLeaderboardParams{
		ID:                  int32(param.ID),
		Name:                param.Name,
		Description:         param.Description,
		AllowAnonymous:      param.AllowAnonymous,
		RequireVerification: param.RequiredVerification,
		UniqueSubmission:    param.UniqueSubmission,
		Descending:          param.Descending,
		CoverImageUrl: pgtype.Text{
			Valid:  true,
			String: param.CoverImageUrl,
		},
	})
	if err != nil {
		return err
	}

	fieldMap := make(map[int]models.Field)
	for _, oldField := range oldLb.Fields {
		fieldMap[oldField.Id] = oldField
	}

	fieldsToAdd := make([]models.Field, 0)
	fieldsToEdit := make([]models.Field, 0)

	for _, field := range param.Fields {
		if field.Id == 0 {
			if field.ForRank {
				return ErrChangeForRank
			}
			fieldsToAdd = append(fieldsToAdd, field)
		} else if oldField, ok := fieldMap[field.Id]; ok {
			if oldField.Type != field.Type {
				return ErrChangeFieldType
			}

			if oldField.ForRank && !field.ForRank || !oldField.ForRank && field.ForRank {
				return ErrChangeFieldType
			}

			if field.ForRank && field.Hidden {
				return ErrHiddenForRank
			}

			if !oldField.Required && field.Required && len(field.Default) == 0 {
				return ErrNoDefault
			}

			if oldField.Name != field.Name ||
				oldField.Hidden != field.Hidden ||
				oldField.Required != field.Required ||
				oldField.FieldOrder != field.FieldOrder {
				fieldsToEdit = append(fieldsToEdit, field)
			}

			if field.Type == string(database.FieldTypeOPTION) {
				diff := arrutil.Diff(oldField.Options, field.Options, func(s1, s2 string) int {
					if s1 != s2 {
						return 1
					}
					return 0
				})

				if len(diff) > 0 {
					err = tx.DeleteLeaderboardOptions(ctx, int32(field.Id))
					if err != nil {
						return err
					}
					param := arrutil.Map(field.Options,
						func(option string) (database.CreateLeaderboardOptionsParams, bool) {
							return database.CreateLeaderboardOptionsParams{
								Fid:    int32(field.Id),
								Option: option,
							}, true
						})
					var n int64
					n, err = tx.CreateLeaderboardOptions(ctx, param)
					if n != int64(len(field.Options)) {
						return ErrUnableToInsertAllOptions
					}
				}

			}

		} else {
			return ErrUnknownEditField
		}
	}

	if len(fieldsToAdd) > 0 {
		var n int64
		n, err = tx.CreateLeaderboardFields(ctx, arrutil.Map(fieldsToAdd, func(field models.Field) (database.CreateLeaderboardFieldsParams, bool) {
			return database.CreateLeaderboardFieldsParams{
				Lid:        int32(param.ID),
				FieldName:  field.Name,
				FieldValue: database.FieldType(field.Type),
				FieldOrder: int32(field.FieldOrder),
				ForRank:    field.ForRank,
				Required:   field.Required,
				Hidden:     field.Hidden,
			}, true
		}))

		if n != int64(len(fieldsToAdd)) {
			return ErrUnableToInsertAllFields
		}
	}

	if len(fieldsToEdit) > 0 {
		type Update struct {
			Id         int
			FieldName  string
			FieldOrder int
			Hidden     bool
			Required   bool
		}
		var data json.RawMessage
		data, err = json.Marshal(arrutil.Map(fieldsToAdd, func(f models.Field) (Update, bool) {
			return Update{
				Id:         f.Id,
				FieldName:  f.Name,
				FieldOrder: f.FieldOrder,
				Hidden:     f.Hidden,
				Required:   f.Required,
			}, true
		}))
		if err != nil {
			return err
		}
		err = tx.BulkUpdateFields(ctx, data)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.ID)
	s.cache.Delete(cacheKeyLBFull)

	return nil
}
