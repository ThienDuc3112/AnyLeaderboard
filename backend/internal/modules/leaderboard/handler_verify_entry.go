package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (s leaderboardService) verifyEntryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("verifyEntryHandler", err) }()

	eidStr := r.PathValue(c.PathValueEntryId)
	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid entry id")
		return
	}

	user, ok := r.Context().Value(c.MiddlewareKeyUser).(database.User)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("context does not give user type")
		return
	}
	lb, ok := r.Context().Value(c.MiddlewareKeyLeaderboard).(database.Leaderboard)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("context does not give leaderboard type")
		return
	}

	body, err := utils.ExtractBody[verifyEntryReqBody](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	err = s.verifyEntry(r.Context(), verifyEntryParam{
		leaderboardId: lb.ID,
		userId:        user.ID,
		entryId:       int32(eid),
		verifyState:   body.Verify,
	})
	if err == errNoEntry {
		utils.RespondWithError(w, 404, "Leaderboard don't have such entry id")
		return
	} else if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	utils.RespondEmpty(w)
}
