package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// TODO
// - Default sorting: by recent no user
// - Add support for sorted by:
//   - author
//   - entries count
func (s leaderboardService) getLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
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

	lbs, err := s.repo.GetRecentLeaderboards(r.Context(), database.GetRecentLeaderboardsParams{
		CreatedAt: pgtype.Timestamptz{Time: cursor.UTC(), Valid: true},
		Limit:     int32(pageSize + 1),
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
		// if true {
		newUrl, _ := url.Parse(r.RequestURI)
		newQuery := newUrl.Query()

		firstLb := lbs[0]
		secondLastLb := lbs[len(lbs)-2]
		log.Printf(
			"\nQuery cursor: %v\t\tunix: %v\nFirst created: %v\tunix: %v\nTime created: %v\tunix: %v\n",
			cursor,
			cursor.UnixMilli(),
			firstLb.CreatedAt.Time,
			firstLb.CreatedAt.Time.UnixMilli(),
			secondLastLb.CreatedAt.Time,
			secondLastLb.CreatedAt.Time.UnixMilli(),
		)
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
