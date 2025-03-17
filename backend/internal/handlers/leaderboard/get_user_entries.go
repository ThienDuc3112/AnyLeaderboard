package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

func (h LeaderboardHandler) GetUserEntries(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("GetUserEntries", err) }()

	lidStr := r.PathValue(c.PathValueLeaderboardId)
	lid, err := strconv.Atoi(lidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid leaderboard id")
		return
	}

	username := r.PathValue(c.PathValueUsername)
	if username == "" {
		utils.RespondWithError(w, 400, "No username provided in query param")
		return
	}

	cursorStr := r.URL.Query().Get(c.QueryParamCursor)
	pageSizeStr := r.URL.Query().Get(c.QueryParamPageSize)
	pageSize := c.DefaultPageSize
	cursor := math.Inf(1)

	if cursorStr != "" {
		curCursor, err := strconv.ParseFloat(cursorStr, 64)
		if err == nil {
			cursor = curCursor
		}
	}

	if pageSizeStr != "" {
		size, err := strconv.Atoi(pageSizeStr)
		if err == nil && size >= c.MinPageSize && size <= c.MaxPageSize {
			pageSize = size
		}
	}

	val, err := h.s.GetEntriesByUser(r.Context(), leaderboard.GetEntriesByUserParam{
		LeaderboardId: int32(lid),
		Username:      username,
		PageSize:      int32(pageSize) + 1,
		Cursor:        cursor,
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

	temp := val.Data
	limit := min(pageSize, len(val.Data))
	val.Data = temp[:limit]
	response := map[string]any{
		"data": val,
	}

	if len(temp) > pageSize {
		newUrl, _ := url.Parse(r.RequestURI)
		newQuery := newUrl.Query()

		secondLastEntry := temp[pageSize-1]
		newQuery.Set(c.QueryParamCursor, fmt.Sprintf("%f", secondLastEntry.SortedField))
		newUrl.RawQuery = newQuery.Encode()
		newUrl.Host = r.Host
		newUrl.Scheme = "https"
		response["next"] = newUrl.String()
	} else {
		response["next"] = nil
	}

	utils.RespondWithJSON(w, 200, response)
}
