package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"net/http"
	"strconv"
)

func (h LeaderboardHandler) GetLeaderboardConfig(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("getLeaderboardHandlerConfig", err) }()

	lidStr := r.PathValue(c.PathValueLeaderboardId)
	lid, err := strconv.Atoi(lidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid leaderboard id")
		return
	}

	res, err := h.s.GetLeaderboard(r.Context(), int32(lid))
	if err != nil {
		switch err {
		case leaderboard.ErrNoLeaderboard:
			utils.RespondWithError(w, 404, "Leaderboard not found")
			err = nil
		default:
			utils.RespondWithError(w, 500, "Internal server error")
		}
		return
	}

	utils.RespondWithJSON(w, 200, res)
}
