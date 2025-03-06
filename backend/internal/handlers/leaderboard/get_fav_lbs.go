package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (h LeaderboardHandler) GetFavoriteLeaderboards(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("GetFavoriteLeaderboards", err) }()

	user, ok := r.Context().Value(c.MiddlewareKeyUser).(database.User)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("context does not give user type")
		return
	}

	cursorStr := r.URL.Query().Get(c.QueryParamCursor)
	pageSizeStr := r.URL.Query().Get(c.QueryParamPageSize)
	pageSize := c.DefaultPageSize
	cursor := time.Now()

	if cursorStr != "" {
		msec, err := strconv.ParseInt(cursorStr, 10, 64)
		if err == nil {
			cursor = time.UnixMilli(msec)
		}
	}

	if pageSizeStr != "" {
		size, err := strconv.Atoi(pageSizeStr)
		if err == nil && size >= c.MinPageSize && size <= c.MaxPageSize {
			pageSize = size
		}
	}

	lbs, err := h.s.GetFavoriteLeaderboards(r.Context(), leaderboard.GetFavLBParams{
		UserID:   user.ID,
		Cursor:   cursor,
		PageSize: int32(pageSize + 1),
	})
	if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	if len(lbs) == 0 {
		utils.RespondWithError(w, 404, "No favorited leaderboards found")
		return
	}
	response := map[string]any{
		"data": lbs[:len(lbs)-1],
	}

	if len(lbs) > pageSize {
		newUrl, _ := url.Parse(r.RequestURI)
		newQuery := newUrl.Query()

		secondLastLb := lbs[len(lbs)-2]
		newQuery.Set(c.QueryParamCursor, fmt.Sprintf("%d", secondLastLb.CreatedAt.UnixMilli()))
		newUrl.RawQuery = newQuery.Encode()
		newUrl.Host = r.Host
		newUrl.Scheme = "https"
		response["next"] = newUrl.String()
	} else {
		response["next"] = nil
	}

	utils.RespondWithJSON(w, 200, response)
}
