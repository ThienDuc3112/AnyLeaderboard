package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
)

func (s leaderboardService) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("createEntryHandler", err) }()

	body, err := utils.ExtractBody[map[string]any](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	lb, ok := r.Context().Value(c.MiddlewareKeyLeaderboard).(database.Leaderboard)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("context does not give leaderboard type")
		return
	}

	var user *database.User
	userData, ok := r.Context().Value(c.MiddlewareKeyUser).(database.User)
	if !ok {
		if !lb.AllowAnnonymous {
			utils.RespondWithError(w, 401, "You are not logged in")
			return
		} else {
			user = nil
		}
	} else {
		user = &userData
	}

	displayName := ""
	displayNameData, exist := body[c.EntryDisplayNameField]
	if exist {
		dn, ok := displayNameData.(string)
		if ok {
			displayName = dn
		}
	}

	eid, fieldName, err := s.createEntry(r.Context(), createEntryParam{
		Leaderboard: lb,
		User:        user,
		Entry:       body,
		DisplayName: displayName,
	})

	if err != nil {
		// Error handling here
		switch err {
		case errNonAnonymousLeaderboard:
			utils.RespondWithError(w, 500, "Internal server error")
			err = fmt.Errorf("no user on nonAnon lb, should've been blocked by middleware")
		case errNoDisplayName:
			utils.RespondWithError(w, 400, "No display name in request or user token found")
			err = nil
		case errRequiredFieldNotExist:
			utils.RespondWithError(w, 400, fmt.Sprintf("field '%s' missing", fieldName))
			err = nil
		case errConflictForRankField:
			utils.RespondWithError(w, 500, fmt.Sprintf("Leaderboard have conflicting field, contact leaderboard owner to resolve '%s' field", fieldName))
			err = fmt.Errorf("field '%s' conflicting for rank: %v", fieldName, err)
		case errOptionFieldNoOptions:
			utils.RespondWithError(w, 500, fmt.Sprintf("Leaderboard have emtpy option field, contact leaderboard owner to resolve '%s' field", fieldName))
			err = fmt.Errorf("field '%s' have no options: %v", fieldName, err)
		case errNotAnOption:
			utils.RespondWithError(w, 400, fmt.Sprintf("field '%s' is not a valid option", fieldName))
			err = nil
		case errUnrankableFieldType:
			utils.RespondWithError(w, 500, fmt.Sprintf("Leaderboard have unrankable field ranked, contact leaderboard owner to resolve '%s' field", fieldName))
			err = fmt.Errorf("field '%s' ranked despite unrankable: %v", fieldName, err)
		case errUnrecognizedField:
			utils.RespondWithError(w, 500, fmt.Sprintf("Leaderboard have unknown field, contact leaderboard owner to resolve '%s' field", fieldName))
			err = fmt.Errorf("field '%s' with unknown/unimplemented field type: %v", fieldName, err)
		default:
			utils.RespondWithError(w, 500, "Internal server error")
		}
		return
	}

	utils.RespondWithJSON(w, 201, map[string]any{
		"id": eid.ID,
	})
}
