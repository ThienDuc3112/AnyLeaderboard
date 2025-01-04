package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
)

func (s leaderboardService) createLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("createLeaderboardHandler", err) }()

	body, err := utils.ExtractBody[createLeaderboardReqBody](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	if err = validate.Struct(body); err != nil {
		utils.RespondToInvalidBody(w, err, trans)
		return
	}

	userCtx := r.Context().Value(c.MiddlewareKeyUser)
	var user database.User
	var ok bool
	if userCtx == nil {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("user context don't exist on a Force Auth path")
		return
	}
	if user, ok = userCtx.(database.User); !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("user context is not of type database.User")
		return
	}

	leaderboard, err := s.createLeaderboard(r.Context(), createLeaderboardParam{
		createLeaderboardReqBody: body,
		User:                     user,
	})

	if err != nil {
		switch err {
		case errMultipleForRankField:
			utils.RespondWithError(w, 400, "Multiple 'For rank' field, only one field can be 'For rank'")
			err = nil
		case errForRankNotRequired:
			utils.RespondWithError(w, 400, "A 'For rank' field must be required")
			err = nil
		case errNoForRankField:
			utils.RespondWithError(w, 400, "No 'For rank' field, one field must be 'For rank'")
			err = nil
		case errNoPublicField:
			utils.RespondWithError(w, 400, "No public field, one field must be not hidden")
			err = nil
		case errNoOptions:
			utils.RespondWithError(w, 400, "An Option field must have atleast one option")
			err = nil
		default:
			utils.RespondWithError(w, 500, "Internal server error")
		}
		return
	}

	utils.RespondWithJSON(w, 201, map[string]any{
		"id": leaderboard.ID,
	})
}
