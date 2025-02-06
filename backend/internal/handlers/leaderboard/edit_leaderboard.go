package leaderboard

import (
	"anylbapi/internal/models"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"log"
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

	body, err := utils.ExtractBody[editLeaderboardBody](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Cannot decode body")
		return
	}
	log.Printf("%+v\n", body)
	if err = validate.Struct(body); err != nil {
		utils.RespondToInvalidBody(w, err, trans)
		return
	}

	err = h.s.EditLeaderboard(r.Context(), leaderboard.EditLeaderboardParam{
		Lid:          body.Lid,
		OldFieldName: body.OldFieldName,
		NewField:     body.NewField,
		NewFieldName: body.NewFieldName,
		DefaultValue: body.NewDefaultValue,
		OldOption:    body.OldOption,
		NewOption:    body.NewOption,
		Action:       leaderboard.EditLeaderboardAction(body.Action),
	})
	if err != nil {
		switch err {
		// Add field errors
		case leaderboard.ErrCannotAddForRank:
			utils.RespondWithError(w, 400, "Cannot add a \"for\" rank field")
		case leaderboard.ErrConflictType:
			utils.RespondWithError(w, 400, "Default value for a field is not correct of correct specify type")
		case leaderboard.ErrUnrecognizedField:
			utils.RespondWithError(w, 400, "Invalid field type")
		case leaderboard.ErrUnableToInsertAllOptions:
			utils.RespondWithError(w, 500, "Unable to create option field")

		// Field validation
		case leaderboard.ErrNoField:
			utils.RespondWithError(w, 400, "Field don't exist")
		case leaderboard.ErrCannotDeleteForRank:
			utils.RespondWithError(w, 400, "\"For rank\" field cannot be remove")

		// Other operation
		case leaderboard.ErrInvalidAction:
			utils.RespondWithError(w, 400, "Invalid action")

		// Generic error
		default:
			utils.RespondWithError(w, 500, "Internal server error")
		}
		return
	}

	utils.RespondEmpty(w)
	// utils.RespondWithError(w, http.StatusNotImplemented, "Not implemented")
}
