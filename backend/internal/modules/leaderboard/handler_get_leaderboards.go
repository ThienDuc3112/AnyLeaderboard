package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// TODO
// - Default sorting: by recent no user
// - Add support for sorted by:
//   - author
//   - entries count
func (s leaderboardService) getLeaderboardsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("getLeaderboardsHandler", err) }()

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

	lbs, err := s.getRecentLeaderboards(r.Context(), getLeaderboardsParam{
		pageSize: pageSize + 1,
		cursor:   cursor,
	})
	if err != nil {
		utils.RespondWithError(w, 500, "Cannot get leaderboards")
		return
	}

	lbPreviews := make([]map[string]any, 0)

	for i, lb := range lbs {
		if i == pageSize {
			break
		}
		lbPreviews = append(lbPreviews, map[string]any{
			"id":            lb.ID,
			"name":          lb.Name,
			"description":   lb.Description,
			"coverImageUrl": lb.CoverImageUrl.String,
			"entriesCount":  lb.EntriesCount,
		})
	}

	response := map[string]any{
		"data": lbPreviews,
	}

	if len(lbs) > pageSize {
		newUrl, _ := url.Parse(r.RequestURI)
		newQuery := newUrl.Query()

		secondLastLb := lbs[len(lbs)-2]
		newQuery.Set(c.QueryParamCursor, fmt.Sprintf("%d", secondLastLb.CreatedAt.Time.UnixMilli()))
		newUrl.RawQuery = newQuery.Encode()
		newUrl.Host = r.Host
		newUrl.Scheme = "https"
		response["next"] = newUrl.String()
	} else {
		response["next"] = nil
	}

	utils.RespondWithJSON(w, 200, response)
}
