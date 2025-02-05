package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"net/http"
	"strconv"
)

func (h LeaderboardHandler) GetVerifiedEntries(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("getVerifiedEntriesHandler", err) }()

	lidStr := r.PathValue(c.PathValueLeaderboardId)
	lid, err := strconv.Atoi(lidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid leaderboard id")
		return
	}

	offsetStr := r.URL.Query().Get("offset")
	pageSizeStr := r.URL.Query().Get(c.QueryParamPageSize)
	pageSize := c.DefaultPageSize
	offset := 0
	if offsetStr != "" {
		oset, err := strconv.Atoi(offsetStr)
		if err == nil {
			offset = oset
		}
	}
	if pageSizeStr != "" {
		size, err := strconv.Atoi(pageSizeStr)
		if err == nil && size >= c.MinPageSize && size <= c.MaxPageSize {
			pageSize = size
		}
	}
	optionStr := r.URL.Query().Get("option")
	requiredVerification, pending, verifiedState, uniqueSubmission := false, false, false, false
	switch optionStr {
	case "pending":
		pending = true
	case "disqualified":
		requiredVerification = true
		verifiedState = false
	case "verified":
		requiredVerification = true
		verifiedState = true
	case "all":
		requiredVerification = false
	}

	lbWithEntries, err := h.s.GetLeaderboardWithEntry(r.Context(), leaderboard.GetLeaderboardParam{
		Id:                   lid,
		PageSize:             pageSize,
		Offset:               offset,
		RequiredVerification: &requiredVerification,
		ForcedPending:        pending,
		VerifyState:          &verifiedState,
		UniqueSubmission:     &uniqueSubmission,
	})

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

	utils.RespondWithJSON(w, 200, lbWithEntries)
}
