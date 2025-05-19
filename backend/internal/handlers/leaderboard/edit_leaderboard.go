package leaderboard

import (
	"anylbapi/internal/models"
	"anylbapi/internal/utils"
	"net/http"
)

type editLeaderboardBody struct {
	Lid             int32         `json:"lid" validate:"required"`
	OldFieldName    string        `json:"oldFieldName" validate:"required_unless=Action ADD"`
	NewFieldName    string        `json:"newFieldName" validate:"required_if=Action RENAME,max=32,isSafeName"`
	Action          string        `json:"action" validate:"required,oneof=RENAME ADD DELETE ADD_OPTION DELETE_OPTION RENAME_OPTION"`
	OldOption       string        `json:"oldOption" validate:"required_if=Action DELETE_OPTION,required_if=Action RENAME_OPTION"`
	NewOption       string        `json:"newOption" validate:"required_if=Action ADD_OPTION,required_if=Action RENAME_OPTION"`
	NewField        *models.Field `json:"newField" validate:"required_if=Action ADD"`
	NewDefaultValue any           `json:"newDefaultValue"`
}

func (h LeaderboardHandler) EditLeaderboard(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("EditLeaderboardHandler", err) }()

	utils.RespondWithError(w, 501, "Haven't implemented")
}
