package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/models"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// TODO
// - Default sorting: by recent no user
// - Search
// - Add support for sorted by:
//   - Author
//   - Entries count
func (h LeaderboardHandler) GetLeaderboards(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("getLeaderboardsHandler", err) }()

	cursorStr := r.URL.Query().Get(c.QueryParamCursor)
	pageSizeStr := r.URL.Query().Get(c.QueryParamPageSize)
	creatorStr := r.URL.Query().Get(c.QueryParamCreatedBy)
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

	var lbs []models.LeaderboardPreview

	if creatorStr != "user" {
		lbs, err = h.s.GetByUsername(r.Context(), leaderboard.GetByUsernameParam{
			PageSize: pageSize + 1,
			Cursor:   cursor,
			Username: creatorStr,
		})
	} else {
		lbs, err = h.s.GetRecents(r.Context(), leaderboard.GetRecentsParam{
			PageSize: pageSize + 1,
			Cursor:   cursor,
		})
	}
	if err != nil {
		utils.RespondWithError(w, 500, "Cannot get leaderboards")
		return
	}

	response := map[string]any{
		"data": lbs[:pageSize],
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
