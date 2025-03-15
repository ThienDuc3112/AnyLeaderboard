package favorite

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/modules/favorite"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (h FavoriteHandler) AddFavorite(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("AddFavoriteHandler", err) }()

	user, ok := r.Context().Value(constants.MidKeyUser).(database.User)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("user context is not of type database.User")
		return
	}

	lidStr := r.PathValue(constants.PathValueLeaderboardId)
	lid, err := strconv.Atoi(lidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid leaderboard id")
		return
	}

	err = h.s.Create(r.Context(), favorite.CreateParam{
		Uid: user.ID,
		Lid: int32(lid),
	})

	if err != nil {
		utils.RespondWithError(w, 400, "Cannot add favorite, leaderboard likely not exist")
		return
	}

	utils.RespondEmpty(w)
}
