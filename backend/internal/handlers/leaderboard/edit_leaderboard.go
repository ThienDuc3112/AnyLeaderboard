package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"strconv"
)

type EditLeaderboardBody struct {
	// ID                   int                   `json:"id" validate:"required"`
	Name                 string                `json:"name" validate:"required,isLBName"`
	Description          string                `json:"description" validate:"max=256"`
	CoverImageUrl        string                `json:"coverImageUrl" validate:"omitempty,http_url"`
	AllowAnonymous       bool                  `json:"allowAnonymous" validate:"excluded_if=UniqueSubmission true"`
	RequiredVerification bool                  `json:"requiredVerification"`
	UniqueSubmission     bool                  `json:"uniqueSubmission"`
	Descending           bool                  `json:"descending"`
	Fields               []models.Field        `json:"fields" validate:"required,min=1,max=10,unique=Name,unique=FieldOrder,dive"`
	ExternalLinks        []models.ExternalLink `json:"externalLinks" validate:"max=5,unique=DisplayValue,dive"`
}

func (h LeaderboardHandler) EditLeaderboard(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("EditLeaderboardHandler", err) }()

	lidStr := r.PathValue(c.PathValueLeaderboardId)
	lid, err := strconv.Atoi(lidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid leaderboard id")
		return
	}

	body, err := utils.ExtractBody[EditLeaderboardBody](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	user, ok := r.Context().Value(c.MidKeyUser).(database.User)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("context does not give user type")
		return
	}

	err = h.s.UpdateLeaderboard(r.Context(), leaderboard.UpdateLeaderboardParam{
		LeaderboardFull: models.LeaderboardFull{
			ID:                   lid,
			Name:                 body.Name,
			Description:          body.Description,
			CoverImageUrl:        body.CoverImageUrl,
			AllowAnonymous:       body.AllowAnonymous,
			RequiredVerification: body.RequiredVerification,
			UniqueSubmission:     body.UniqueSubmission,
			Descending:           body.Descending,
			ExternalLink:         body.ExternalLinks,
			Fields:               body.Fields,
		},
		UserId: int(user.ID),
	})

	if err != nil {
		switch err {
		case leaderboard.ErrNoLeaderboard:
			utils.RespondWithError(w, 404, "Leaderboard not found")
			err = nil
		case leaderboard.ErrNotOwnLeaderboard:
			utils.RespondWithError(w, 403, "Not own leaderboard")
			err = nil
		case leaderboard.ErrChangeForRank:
			utils.RespondWithError(w, http.StatusBadRequest, "Cannot change for rank field")
			err = nil
		case leaderboard.ErrChangeFieldType:
			utils.RespondWithError(w, http.StatusBadRequest, "Cannot change field type")
			err = nil
		case leaderboard.ErrHiddenForRank:
			utils.RespondWithError(w, http.StatusBadRequest, "Cannot hide for rank field")
			err = nil
		case leaderboard.ErrNoDefault:
			utils.RespondWithError(w, http.StatusBadRequest, "New required field don't have a default value")
			err = nil
		case leaderboard.ErrUnknownEditField:
			utils.RespondWithError(w, http.StatusBadRequest, "Try to edit a non-existent field")
			err = nil
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	utils.RespondEmpty(w)
}
