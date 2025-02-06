package leaderboard

import (
	"anylbapi/internal/models"
	"context"
)

type EditLeaderboardAction string

const (
	addField           EditLeaderboardAction = "ADD"
	renameField        EditLeaderboardAction = "RENAME"
	deleteField        EditLeaderboardAction = "DELETE"
	addOptionsField    EditLeaderboardAction = "ADD_OPTION"
	deleteOptionsField EditLeaderboardAction = "DELETE_OPTION"
	renameOptionsField EditLeaderboardAction = "RENAME_OPTION"
)

type EditLeaderboardParam struct {
	Lid          int32
	OldFieldName string
	NewFieldName string
	NewField     *models.Field
	DefaultValue any
	OldOption    string
	NewOption    string
	Action       EditLeaderboardAction `validate:"oneof=RENAME ADD DELETE ADD_OPTION DELETE_OPTION RENAME_OPTION"`
}

func (s LeaderboardService) EditLeaderboard(ctx context.Context, param EditLeaderboardParam) error {
	if err := validate.Struct(param); err != nil {
		return ErrInvalidAction
	}

	// The only action that doesn't need old field
	if param.Action == addField {
		return s.AddField(ctx, AddFieldParam{
			Lid:          param.Lid,
			NewField:     *param.NewField,
			DefaultValue: param.DefaultValue,
		})
	}

	field, err := s.GetField(ctx, GetFieldParam{
		Lid:       param.Lid,
		FieldName: param.OldFieldName,
	})
	if err != nil {
		return err
	}

	switch param.Action {
	case renameField:
		return s.RenameField(ctx, RenameFieldParams{
			fieldName: param.OldFieldName,
			newName:   param.NewFieldName,
			lid:       param.Lid,
		})
	case deleteField:
		if field.ForRank {
			return ErrCannotDeleteForRank
		}

		return s.DeleteField(ctx, DeleteFieldParam{
			Lid:          param.Lid,
			OldFieldName: param.OldFieldName,
		})
	case addOptionsField:
		return s.AddOptionToField(ctx, AddOptionToFieldParam{
			FieldName: param.OldFieldName,
			Lid:       param.Lid,
			NewOption: param.NewOption,
		})
	case deleteOptionsField:
		return s.DeleteOption(ctx, DeleteOptionParam{
			Lid:       param.Lid,
			FieldName: param.OldFieldName,
			Option:    param.OldOption,
		})
	case renameOptionsField:
		return s.RenameOption(ctx, RenameOptionParam{
			Lid:       param.Lid,
			FieldName: param.OldFieldName,
			OldOption: param.OldOption,
			NewOption: param.NewOption,
		})
	default:
		return ErrInvalidAction
	}
}
