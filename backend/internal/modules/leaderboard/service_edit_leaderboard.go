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
	Lid        int32
	OldFieldId int32
	NewField   struct {
		field
		defaultValue any
	}
	Action editLeaderboardAction `validate:"oneof=RENAME ADD DELETE REPLACE ADD_OPTION DELETE_OPTION REPLACE_OPTION"`
}

func (s leaderboardService) editLeaderboard(ctx context.Context, param editLeaderboardParam) error {
	if err := validate.Struct(param); err != nil {
		return errInvalidAction
	}

	if param.Action == addField {
		return s.addField(ctx, addFieldParam{
			Lid:      param.Lid,
			NewField: param.NewField,
		})
	}

	return fmt.Errorf("unimplemented")
}
