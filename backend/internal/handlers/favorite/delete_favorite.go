package favorite

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/modules/favorite"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (h FavoriteHandler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("AddFavoriteHandler", err) }()

	user, ok := r.Context().Value(c.MiddlewareKeyUser).(database.User)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("user context is not of type database.User")
		return
	}

	lidStr := r.PathValue(c.PathValueLeaderboardId)
	lid, err := strconv.Atoi(lidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid leaderboard id")
		return
	}

	err = h.s.DeleteFavorite(r.Context(), favorite.DeleteParam{
		Uid: user.ID,
		Lid: int32(lid),
	})

	if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	utils.RespondEmpty(w)
}
