package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (h LeaderboardHandler) VerifyEntry(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("verifyEntryHandler", err) }()

	eidStr := r.PathValue(c.PathValueEntryId)
	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid entry id")
		return
	}

	user, ok := r.Context().Value(c.MiddlewareKeyUser).(database.User)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("context does not give user type")
		return
	}
	lb, ok := r.Context().Value(c.MiddlewareKeyLeaderboard).(models.LeaderboardPreview)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("context does not give leaderboard type")
		return
	}

	type verifyEntryReqBody struct {
		Verify bool `json:"verify"`
	}

	body, err := utils.ExtractBody[verifyEntryReqBody](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	err = h.s.VerifyEntry(r.Context(), leaderboard.VerifyEntryParam{
		LeaderboardId: int32(lb.ID),
		UserId:        user.ID,
		EntryId:       int32(eid),
		VerifyState:   body.Verify,
	})
	if err == leaderboard.ErrNoEntry {
		utils.RespondWithError(w, 404, "Leaderboard don't have such entry id")
		return
	} else if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	utils.RespondEmpty(w)
}
