package leaderboard

import (
	"context"
)

type editLeaderboardAction string

const (
	addField           editLeaderboardAction = "ADD"
	renameField        editLeaderboardAction = "RENAME"
	deleteField        editLeaderboardAction = "DELETE"
	replaceField       editLeaderboardAction = "REPLACE"
	addOptionsField    editLeaderboardAction = "ADD_OPTION"
	deleteOptionsField editLeaderboardAction = "DELETE_OPTION"
	renameOptionsField editLeaderboardAction = "RENAME_OPTION"
)

type editLeaderboardParam struct {
	Lid          int32
	OldFieldName string
	NewField     struct {
		field
		defaultValue any
	}
	OldOption string
	NewOption string
	Action    editLeaderboardAction `validate:"oneof=RENAME ADD DELETE REPLACE ADD_OPTION DELETE_OPTION RENAME_OPTION"`
}

func (s leaderboardService) editLeaderboard(ctx context.Context, param editLeaderboardParam) error {
	if err := validate.Struct(param); err != nil {
		return ErrInvalidAction
	}

	// The only action that doesn't need old field
	if param.Action == addField {
		return s.addField(ctx, addFieldParam{
			Lid:      param.Lid,
			NewField: param.NewField,
		})
	}

	field, err := s.getField(ctx, getFieldParam{
		Lid:       param.Lid,
		FieldName: param.OldFieldName,
	})
	if err != nil {
		return err
	}

	switch param.Action {
	case renameField:
		return s.renameField(ctx, renameFieldParams{
			fieldName: param.OldFieldName,
			newName:   param.NewField.Name,
			lid:       param.Lid,
		})
	case deleteField:
		if field.ForRank {
			return ErrCannotDeleteForRank
		}

		return s.deleteField(ctx, deleteFieldParam{
			Lid:          param.Lid,
			OldFieldName: param.OldFieldName,
		})
	case addOptionsField:
		return s.addOptionToField(ctx, addOptionToFieldParam{
			FieldName: param.OldFieldName,
			Lid:       param.Lid,
			NewOption: param.NewOption,
		})
	case deleteOptionsField:
		return s.deleteOption(ctx, deleteOptionParam{
			Lid:       param.Lid,
			FieldName: param.OldFieldName,
			Option:    param.OldOption,
		})
	case renameOptionsField:
		return s.renameOption(ctx, renameOptionParam{
			Lid:       param.Lid,
			FieldName: param.OldFieldName,
			OldOption: param.OldOption,
			NewOption: param.NewOption,
		})
	default:
		return ErrInvalidAction
	}
}
