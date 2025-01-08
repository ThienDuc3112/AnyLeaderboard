package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
)

func (s leaderboardService) addVerifierHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("addVerifierHandler", err) }()

	lb, ok := r.Context().Value(c.MiddlewareKeyLeaderboard).(database.Leaderboard)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("user context is not of type database.Leaderboard")
		return
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

	err = s.addVerifier(r.Context(), addVerifierParam{
		username: body.Username,
		lid:      lb.ID,
	})

	if err == errNoUser {
		utils.RespondWithError(w, 400, "User don't exist")
		return
	} else if err == errAlreadyVerifier {
		utils.RespondWithError(w, 400, "User already a verifier")
		return
	} else if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	utils.RespondEmpty(w)
}
