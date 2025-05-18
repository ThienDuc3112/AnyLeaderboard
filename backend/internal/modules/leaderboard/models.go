package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"errors"
	"time"
)

// ============ Service param and return types ============
type CreateLeaderboardParam struct {
	models.SubmitLBStructure
	User database.User
}

type CreateEntryParam struct {
	Leaderboard models.Leaderboard
	Entry       map[string]any
	User        *database.User
	DisplayName string
}

type DeleteEntryParam struct {
	UserId      int32
	Leaderboard models.Leaderboard
	EntryId     int
}

type GetLeaderboardParam struct {
	Id                   int
	PageSize             int
	Offset               int
	UniqueSubmission     *bool
	RequiredVerification *bool
	VerifyState          *bool
	ForcedPending        bool
}

type GetRecentsParam struct {
	PageSize int
	Cursor   time.Time
}

type AddVerifierParam struct {
	Username string
	Lid      int32
}

type GetEntriesParam struct {
	Lid                  int32
	RequiredVerification bool
	UniqueSubmission     bool
	VerifyState          bool
	ForcedPending        bool
	Desc                 bool
	Offset               int32
	PageSize             int32
}

type GetEntriesReturn struct {
	Entries []database.LeaderboardEntry
	Count   int64
}

type VerifyEntryParam struct {
	LeaderboardId int32
	UserId        int32
	EntryId       int32
	VerifyState   bool
}

// ============ Service errors ============
var (
	ErrMultipleForRankField     = errors.New("multiple for rank field")
	ErrNoForRankField           = errors.New("no for rank field found")
	ErrForRankNotRequired       = errors.New("for rank need to be required")
	ErrNoPublicField            = errors.New("no public field exist")
	ErrNoOptions                = errors.New("option field have no options")
	ErrUnableToInsertAllFields  = errors.New("unable to insert all fields")
	ErrUnableToInsertAllOptions = errors.New("unable to insert all options")
	ErrUnableToInsertAllLinks   = errors.New("unable to insert all links")

	ErrRequiredFieldNotExist   = errors.New("required field must be filled in")
	ErrUnrecognizedField       = errors.New("unrecognized field type")
	ErrOptionFieldNoOptions    = errors.New("option field don't have any options")
	ErrNotAnOption             = errors.New("value not in option")
	ErrNonAnonymousLeaderboard = errors.New("leaderboard required user account")
	ErrConflictForRankField    = errors.New("more than 1 for rank field found")
	ErrUnrankableFieldType     = errors.New("field type cannot be rank")
	ErrNoDisplayName           = errors.New("no user or display name found")

	ErrNoLeaderboard = errors.New("leaderboard don't exist")

	ErrNoEntry       = errors.New("entry don't exist")
	ErrNotAuthorized = errors.New("not authorized to perform such action")

	ErrNoUser          = errors.New("user don't exist")
	ErrAlreadyVerifier = errors.New("user already a verifier")

	ErrInvalidAction    = errors.New("action is invalid")
	ErrCannotAddForRank = errors.New("cannot add for rank field")
	ErrConflictType     = errors.New("default value not of correct type to field type")

	ErrNoField = errors.New("field don't exist")

	ErrCannotDeleteForRank = errors.New("cannot delete for rank field")
)
