package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
)

func (h LeaderboardHandler) getVerifiers(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("getVerifiersHandler", err) }()

	lb, ok := r.Context().Value(c.MiddlewareKeyLeaderboard).(database.Leaderboard)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("user context is not of type database.Leaderboard")
		return
	}

	verifiers, err := h.s.GetVerifiers(r.Context(), lb.ID)
	if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	res := make([]map[string]any, 0)

	for _, verifier := range verifiers {
		newVerifier := map[string]any{
			"username":    verifier.Username,
			"displayName": verifier.DisplayName,
			"description": verifier.Description,
			"createdAt":   verifier.CreatedAt.Time,
		}
		res = append(res, newVerifier)
	}

	utils.RespondWithJSON(w, 200, res)
}
