package leaderboard

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/middleware"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
)

func (s leaderboardService) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("createEntryHandler", err) }()

	body, err := utils.ExtractBody[map[string]any](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	lb, ok := r.Context().Value(middleware.KeyLeaderboard).(database.Leaderboard)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("context does not give leaderboard type")
		return
	}

	var user *database.User
	userData, ok := r.Context().Value(middleware.KeyUser).(database.User)
	if !ok {
		if lb.RequireVerification {
			utils.RespondWithError(w, 500, "Internal server error")
			err = fmt.Errorf("context does not give user type")
			return
		} else {
			user = nil
		}
	} else {
		user = &userData
	}

	displayName := ""
	displayNameData, ok := body[constants.EntryDisplayNameField]
	if ok {
		dn, ok := displayNameData.(string)
		if ok {
			displayName = dn
		}
	}

	eid, err, fieldName := s.createEntry(r.Context(), createEntryParam{
		Leaderboard: lb,
		User:        user,
		Entry:       body,
		DisplayName: displayName,
	})

	if err != nil {
		// Error handling here
		switch err {

		}
		return
	}

	utils.RespondWithJSON(w, 201, map[string]any{
		"id": eid.ID,
	})
}
