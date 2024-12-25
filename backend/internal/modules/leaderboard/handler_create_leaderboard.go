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
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	utils.RespondWithJSON(w, 201, map[string]any{
		"id": leaderboard.ID,
	})
}
