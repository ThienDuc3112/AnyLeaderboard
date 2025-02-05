package leaderboard

import (
	"anylbapi/internal/database"
	"encoding/json"
	"errors"
	"time"
)

// ============ Request body type ============
type createLeaderboardReqBody struct {
	Name                 string         `json:"name" validate:"required,isLBName"`
	Description          string         `json:"description" validate:"max=256"`
	CoverImageURL        string         `json:"coverImageUrl" validate:"omitempty,http_url"`
	ExternalLinks        []externalLink `json:"externalLinks" validate:"max=5,unique=DisplayValue,dive"`
	AllowAnonymous       bool           `json:"allowAnonymous"`
	RequiredVerification bool           `json:"requiredVerification"`
	UniqueSubmission     bool           `json:"uniqueSubmission" validate:"excluded_if=RequiredVerification false"`
	Fields               []field        `json:"fields" validate:"required,min=1,max=10,unique=Name,unique=FieldOrder,dive"`
}

type entry struct {
	Id         int             `json:"id"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
	Fields     json.RawMessage `json:"fields"`
	Verified   bool            `json:"verified"`
	VerifiedAt *time.Time      `json:"verifiedAt,omitempty"`
	VerifiedBy string          `json:"verifiedBy,omitempty"`
}

type externalLink struct {
	DisplayValue string `json:"displayValue" validate:"required,max=32"`
	URL          string `json:"url" validate:"required,http_url"`
}

type field struct {
	Name       string   `json:"name" validate:"required,max=32,isSafeName"`
	Required   bool     `json:"required"`
	Hidden     bool     `json:"hidden"`
	FieldOrder int      `json:"fieldOrder" validate:"required"`
	Type       string   `json:"type" validate:"required,oneof=TEXT NUMBER DURATION TIMESTAMP OPTION"`
	Options    []string `json:"options" validate:"required_if=Type OPTION,omitempty,unique,min=1,dive,min=1,max=32,isSafeName"`
	ForRank    bool     `json:"forRank" validate:"excluded_if=Type OPTION"`
}

type addVerifierReqBody struct {
	Username string `json:"username" validate:"required,min=3,max=64,isUsername"`
}

type verifyEntryReqBody struct {
	Verify bool `json:"verify"`
}

// ============ Service param and return types ============
type createLeaderboardParam struct {
	createLeaderboardReqBody
	User database.User
}

type createEntryParam struct {
	Leaderboard database.Leaderboard
	Entry       map[string]interface{}
	User        *database.User
	DisplayName string
}

type deleteEntryParam struct {
	user        database.User
	leaderboard database.Leaderboard
	entryId     int
}

type getLeaderboardParam struct {
	id                   int
	pageSize             int
	offset               int
	uniqueSubmission     *bool
	requiredVerification *bool
	verifyState          *bool
	forcedPending        bool
}

type getLeaderboardsParam struct {
	pageSize int
	cursor   time.Time
}

type leaderboardWithEntry struct {
	ID                   int            `json:"id"`
	Name                 string         `json:"name"`
	Description          string         `json:"description"`
	CoverImageUrl        string         `json:"coverImageUrl,omitempty"`
	EntriesCount         int            `json:"entriesCount"`
	AllowAnonymous       bool           `json:"allowAnonymous"`
	RequiredVerification bool           `json:"requiredVerification"`
	UniqueSubmission     bool           `json:"uniqueSubmission"`
	ExternalLink         []externalLink `json:"externalLinks"`
	Fields               []field        `json:"fields"`
	Data                 []entry        `json:"data"`
}

type addVerifierParam struct {
	username string
	lid      int32
}

type getEntriesParam struct {
	lid                  int32
	RequiredVerification bool
	UniqueSubmission     bool
	VerifyState          bool
	ForcedPending        bool
	offset               int32
	pageSize             int32
}

type getEntriesReturn struct {
	entries []database.LeaderboardEntry
	count   int64
}

type verifyEntryParam struct {
	leaderboardId int32
	userId        int32
	entryId       int32
	verifyState   bool
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
