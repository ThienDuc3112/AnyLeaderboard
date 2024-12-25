package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/middleware"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
)

func (s leaderboardService) createLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("signupHandler", err) }()

	body, err := utils.ExtractBody[createLeaderboardReqBody](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	if err = validate.Struct(body); err != nil {
		utils.RespondToInvalidBody(w, err, trans)
		return
	}
	if len(body.Fields) == 0 {
		utils.RespondWithError(w, 400, "There must be atleast a field")
		return
	}

	userCtx := r.Context().Value(middleware.KeyUser)
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
		case errNoForRankField:
			utils.RespondWithError(w, 400, "No 'For rank' field, one field must be 'For rank'")
		case errNoPublicField:
			utils.RespondWithError(w, 400, "No public field, one field must be not hidden")
		default:
			utils.RespondWithError(w, 500, "Internal server error")
		}
		return
	}

	utils.RespondWithJSON(w, 201, map[string]any{
		"id": leaderboard.ID,
	})
}
