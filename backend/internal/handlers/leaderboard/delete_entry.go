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

func (h LeaderboardHandler) DeleteEntry(w http.ResponseWriter, r *http.Request) {
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

	lb, ok := r.Context().Value(c.MidKeyLeaderboard).(models.Leaderboard)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("context does not give leaderboard type")
		return
	}

	err = h.s.DeleteEntry(r.Context(), leaderboard.DeleteEntryParam{UserId: user.ID, Leaderboard: lb, EntryId: eid})

	if err != nil {
		switch err {
		case leaderboard.ErrNoEntry:
			utils.RespondWithError(w, 404, "Entry not found")
			err = nil
		case leaderboard.ErrNotAuthorized:
			utils.RespondWithError(w, 403, "No permission to delete this")
			err = nil
		default:
			utils.RespondWithError(w, 500, "Internal server error")
		}
		return
	}

	utils.RespondEmpty(w)
}
