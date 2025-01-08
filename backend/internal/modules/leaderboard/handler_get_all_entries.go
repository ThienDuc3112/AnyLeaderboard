package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/utils"
	"net/http"
	"strconv"
)

func (s leaderboardService) getAllEntriesHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("getAllEntriesHandler", err) }()

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

	// How is this allowed in go???
	true := true

	lbWithEntries, err := s.getLeaderboardWithEntry(r.Context(), getLeaderboardParam{
		id:                   lid,
		pageSize:             pageSize,
		offset:               offset,
		requiredVerification: &true,
	})

	if err != nil {
		switch err {
		case errNoLeaderboard:
			utils.RespondWithError(w, 404, "Leaderboard not found")
			err = nil
		default:
			utils.RespondWithError(w, 500, "Internal server error")
		}
		return
	}

	utils.RespondWithJSON(w, 200, lbWithEntries)
}
