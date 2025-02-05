package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
)

func (h LeaderboardHandler) RemoveVerifier(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("removeVerifiersHandler", err) }()

	lb, ok := r.Context().Value(c.MiddlewareKeyLeaderboard).(database.Leaderboard)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("user context is not of type database.Leaderboard")
		return
	}

	type addVerifierReqBody struct {
		Username string `json:"username" validate:"required,min=3,max=64,isUsername"`
	}

	body, err := utils.ExtractBody[addVerifierReqBody](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	if err = validate.Struct(body); err != nil {
		utils.RespondToInvalidBody(w, err, trans)
		return
	}

	err = h.s.RemoveVerifier(r.Context(), leaderboard.AddVerifierParam{
		Username: body.Username,
		Lid:      lb.ID,
	})

	if err == leaderboard.ErrNoUser {
		utils.RespondWithError(w, 400, "User don't exist")
		return
	} else if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	utils.RespondEmpty(w)
}
