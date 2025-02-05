package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (s leaderboardService) deleteEntryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("deleteEntryHandler", err) }()

	eidStr := r.PathValue(c.PathValueEntryId)
	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid entry id")
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

	err = s.deleteEntry(r.Context(), deleteEntryParam{user: user, leaderboard: lb, entryId: eid})

	if err != nil {
		switch err {
		case ErrNoEntry:
			utils.RespondWithError(w, 404, "Entry not found")
			err = nil
		case ErrNotAuthorized:
			utils.RespondWithError(w, 403, "No permission to delete this")
			err = nil
		default:
			utils.RespondWithError(w, 500, "Internal server error")
		}
		return
	}

	utils.RespondEmpty(w)
}
