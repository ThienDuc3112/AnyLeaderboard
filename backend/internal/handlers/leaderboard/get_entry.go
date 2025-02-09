package leaderboard

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"net/http"
	"strconv"
)

func (h LeaderboardHandler) GetEntry(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("getEntry", err) }()

	lidStr := r.PathValue(constants.PathValueLeaderboardId)
	lid, err := strconv.Atoi(lidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid leaderboard id")
		return
	}

	eidStr := r.PathValue(constants.PathValueEntryId)
	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid entry id")
		return
	}

	val, err := h.s.GetEntry(r.Context(), leaderboard.GetEntryParam{
		Eid: int32(eid),
		Lid: int32(lid),
	})

	if err != nil {
		switch err {
		case leaderboard.ErrNoEntry:
			utils.RespondWithError(w, 404, "Entry don't exist")
			err = nil
		case leaderboard.ErrNoLeaderboard:
			utils.RespondWithError(w, 404, "Leaderboard don't exist")
			err = nil
		default:
			utils.RespondWithError(w, 500, "Internal server error")
		}
		return
	}

	utils.RespondWithJSON(w, 200, map[string]any{
		"leaderboard": val.Leaderboard,
		"entry":       val.Entry,
	})
}
