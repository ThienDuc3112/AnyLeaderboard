package leaderboard

import (
	"anylbapi/internal/models"
	"context"
)

type editLeaderboardAction string

const (
	addField           editLeaderboardAction = "ADD"
	renameField        editLeaderboardAction = "RENAME"
	deleteField        editLeaderboardAction = "DELETE"
	addOptionsField    editLeaderboardAction = "ADD_OPTION"
	deleteOptionsField editLeaderboardAction = "DELETE_OPTION"
	renameOptionsField editLeaderboardAction = "RENAME_OPTION"
)

type EditLeaderboardParam struct {
	Lid          int32
	OldFieldName string
	NewField     struct {
		models.Field
		defaultValue any
	}
	OldOption string
	NewOption string
	Action    editLeaderboardAction `validate:"oneof=RENAME ADD DELETE ADD_OPTION DELETE_OPTION RENAME_OPTION"`
}

func (s LeaderboardService) EditLeaderboard(ctx context.Context, param EditLeaderboardParam) error {
	if err := validate.Struct(param); err != nil {
		return ErrInvalidAction
	}

	// The only action that doesn't need old field
	if param.Action == addField {
		return s.AddField(ctx, AddFieldParam{
			Lid:      param.Lid,
			NewField: param.NewField,
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
			newName:   param.NewField.Name,
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
