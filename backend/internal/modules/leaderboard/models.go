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
	VerifiedAt time.Time       `json:"verifiedAt,omitempty"`
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
	id       int
	pageSize int
	offset   int
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

// ============ Service errors ============
var (
	errMultipleForRankField     = errors.New("multiple for rank field")
	errNoForRankField           = errors.New("no for rank field found")
	errForRankNotRequired       = errors.New("for rank need to be required")
	errNoPublicField            = errors.New("no public field exist")
	errNoOptions                = errors.New("option field have no options")
	errUnableToInsertAllFields  = errors.New("unable to insert all fields")
	errUnableToInsertAllOptions = errors.New("unable to insert all options")
	errUnableToInsertAllLinks   = errors.New("unable to insert all links")

	errRequiredFieldNotExist   = errors.New("required field must be filled in")
	errUnrecognizedField       = errors.New("unrecognized field type")
	errOptionFieldNoOptions    = errors.New("option field don't have any options")
	errNotAnOption             = errors.New("value not in option")
	errNonAnonymousLeaderboard = errors.New("leaderboard required user account")
	errConflictForRankField    = errors.New("more than 1 for rank field found")
	errUnrankableFieldType     = errors.New("field type cannot be rank")
	errNoDisplayName           = errors.New("no user or display name found")

	errNoLeaderboard = errors.New("leaderboard don't exist")

	errNoEntry       = errors.New("entry don't exist")
	errNotAuthorized = errors.New("not authorized to perform such action")
)
