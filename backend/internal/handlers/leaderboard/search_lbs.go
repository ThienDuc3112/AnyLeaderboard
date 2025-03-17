package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

func (h LeaderboardHandler) Search(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("GetFavoriteLeaderboards", err) }()

	user, _ := r.Context().Value(c.MidKeyUser).(database.User)

	cursorStr := r.URL.Query().Get(c.QueryParamCursor)
	pageSizeStr := r.URL.Query().Get(c.QueryParamPageSize)
	SearchStr := r.URL.Query().Get(c.QueryParamSearch)
	pageSize := c.DefaultPageSize
	cursor := float32(math.Inf(1))

	if SearchStr == "" {
		utils.RespondWithError(w, 400, fmt.Sprintf("Query param \"%s\" must be present", c.QueryParamSearch))
		return
	}

	if cursorStr != "" {
		num, err := strconv.ParseFloat(cursorStr, 32)
		if err == nil {
			cursor = float32(num)
		}
	}

	if pageSizeStr != "" {
		size, err := strconv.Atoi(pageSizeStr)
		if err == nil && size >= c.MinPageSize && size <= c.MaxPageSize {
			pageSize = size
		}
	}

	res, err := h.s.Search(r.Context(), leaderboard.SearchParam{
		Term:         SearchStr,
		UserId:       user.ID,
		PageSize:     int32(pageSize + 1),
		SearchCursor: cursor,
	})
	if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	lbs := make([]models.LeaderboardPreview, len(res.Leaderboards))
	for i, row := range res.Leaderboards {
		lbs[i] = models.LeaderboardPreview{
			ID:             row.ID,
			Name:           row.Name,
			Description:    row.Description,
			CoverImageUrl:  row.CoverImageUrl,
			EntriesCount:   res.EntryCounts[i],
			CreatedAt:      row.CreatedAt,
			AllowAnonymous: row.AllowAnonymous,
		}
	}

	if len(lbs) == 0 {
		utils.RespondWithError(w, 404, "No favorited leaderboards found")
		return
	}
	response := map[string]any{
		"data": lbs[:min(pageSize, len(lbs))],
	}

	if len(lbs) > pageSize && res.Rank[len(res.Rank)-1] > 0 {
		newUrl, _ := url.Parse(r.RequestURI)
		newQuery := newUrl.Query()

		secondLastLb := res.Rank[len(lbs)-2]
		newQuery.Set(c.QueryParamCursor, fmt.Sprintf("%f", secondLastLb-0.000001))
		newUrl.RawQuery = newQuery.Encode()
		newUrl.Host = r.Host
		newUrl.Scheme = "https"
		response["next"] = newUrl.String()
	} else {
		response["next"] = nil
	}

	utils.RespondWithJSON(w, 200, response)

}
