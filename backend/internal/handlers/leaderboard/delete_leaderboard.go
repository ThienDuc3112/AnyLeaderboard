package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (h LeaderboardHandler) DeleteLeaderboard(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("deleteEntryHandler", err) }()

	lidStr := r.PathValue(c.PathValueLeaderboardId)
	lid, err := strconv.Atoi(lidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid leaderboard id")
		return
	}

	user, ok := r.Context().Value(c.MidKeyUser).(database.User)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("context does not give user type")
		return
	}

	err = h.s.DeleteLeaderboard(r.Context(), leaderboard.DeleteLeaderboardParam{
		UserID:        int(user.ID),
		LeaderboardID: lid,
	})

	if err != nil {
		switch err {
		case leaderboard.ErrNotOwnLeaderboard:
			utils.RespondWithError(w, http.StatusForbidden, "Not own leaderboard")
			err = nil
		case leaderboard.ErrNoLeaderboard:
			utils.RespondWithError(w, http.StatusForbidden, "Not own leaderboard")
			err = nil
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	utils.RespondEmpty(w)
}
