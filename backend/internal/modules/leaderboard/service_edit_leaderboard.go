package leaderboard

import (
	"context"
	"fmt"
)

type editLeaderboardAction string

const (
	addField            editLeaderboardAction = "ADD"
	renameField         editLeaderboardAction = "RENAME"
	deleteField         editLeaderboardAction = "DELETE"
	replaceField        editLeaderboardAction = "REPLACE"
	addOptionsField     editLeaderboardAction = "ADD_OPTION"
	deleteOptionsField  editLeaderboardAction = "DELETE_OPTION"
	replaceOptionsField editLeaderboardAction = "REPLACE_OPTION"
)

type editLeaderboardParam struct {
	Lid          int32
	OldFieldName string
	NewField     struct {
		field
		defaultValue any
	}
	NewOption string
	Action    editLeaderboardAction `validate:"oneof=RENAME ADD DELETE REPLACE ADD_OPTION DELETE_OPTION REPLACE_OPTION"`
}

func (s leaderboardService) editLeaderboard(ctx context.Context, param editLeaderboardParam) error {
	if err := validate.Struct(param); err != nil {
		return errInvalidAction
	}

	switch param.Action {
	case addField:
		return s.addField(ctx, addFieldParam{
			Lid:      param.Lid,
			NewField: param.NewField,
		})
	case renameField:
		_, err := s.getField(ctx, getFieldParam{
			Lid:       param.Lid,
			FieldName: param.OldFieldName,
		})
		if err != nil {
			return err
		}

		return s.renameField(ctx, renameFieldParams{
			fieldName: param.OldFieldName,
			newName:   param.NewField.Name,
			lid:       param.Lid,
		})
	case deleteField:
		field, err := s.getField(ctx, getFieldParam{
			Lid:       param.Lid,
			FieldName: param.OldFieldName,
		})
		if err != nil {
			return err
		}
		if field.ForRank {
			return errCannotDeleteForRank
		}

		return s.deleteField(ctx, deleteFieldParam{
			Lid:          param.Lid,
			OldFieldName: param.OldFieldName,
		})
	case addOptionsField:
		_, err := s.getField(ctx, getFieldParam{
			Lid:       param.Lid,
			FieldName: param.OldFieldName,
		})
		if err != nil {
			return err
		}

		return s.addOptionToField(ctx, addOptionToFieldParam{
			FieldName: param.OldFieldName,
			Lid:       param.Lid,
			NewOption: param.NewOption,
		})
	case deleteOptionsField:
	case replaceOptionsField:
	default:
		// return fmt.Errorf("unknown action")
		return errInvalidAction
	}
	return fmt.Errorf("unimplemented")
}
